package ch.bfh.ti.hirtp1ganzg1.thesis.api.views

import Signature
import ch.bfh.ti.hirtp1ganzg1.thesis.api.marshalling.*
import ch.bfh.ti.hirtp1ganzg1.thesis.api.services.*
import ch.bfh.ti.hirtp1ganzg1.thesis.api.utils.*
import com.google.protobuf.ByteString
import io.ktor.application.call
import io.ktor.locations.KtorExperimentalLocationsAPI
import io.ktor.locations.locations
import io.ktor.request.receive
import io.ktor.response.respond
import io.ktor.routing.Routing
import io.ktor.routing.post
import org.koin.ktor.ext.inject
import java.io.File

@KtorExperimentalLocationsAPI
fun Routing.sign() {
    val oidcService by inject<IOIDCService>()
    val secretService by inject<ISecretService>()
    val signingKeyService by inject<ISigningKeysService>()
    val caService by inject<ICertificateAuthorityService>()
    val tsaService by inject<ITimestampingService>()
    val signatureHoldingService by inject<ISignaturesHoldingService>()

    // TODO refactor this monster into smaller methods
    post(URLs.SIGN) {
        when (val input = call.receive<SigningRequest>().validate()) {
            is Valid -> {
                val sortedHashes = input.value.hashes.sorted()
                val salt = calculateSalt(
                    secretService.hkdf(hexStringToByteArray(input.value.seed)),
                    sortedHashes
                )
                if (byteArrayToHexString(salt) != input.value.salt) {
                    throw InvalidDataException("Salt mismatch")
                }

                val jwtValidationResult = oidcService.validateIdToken(input.value.id_token)
                val maskedHashes = maskHashes(sortedHashes, salt)
                if (calculateOidcNonce(maskedHashes.concatenate()) != jwtValidationResult.idToken.getClaim("nonce").asString()) {
                    throw InvalidDataException("Nonce mismatch")
                }

                when (val subjectInformation = SigningKeySubjectInformation.fromIdToken(jwtValidationResult.idToken)) {
                    is Either.Success -> {
                        caService.signCSR(
                            signingKeyService.generateSigningKey(subjectInformation.value)
                        ).also { certificateHolder ->
                            signingKeyService.signToPkcs7(
                                subjectInformation.value,
                                Signature.SignatureData.newBuilder()
                                    .addAllSaltedDocumentHash(
                                        maskedHashes.map { it.toByteString() }
                                    )
                                    .setHashAlgorithm(Signature.HashAlgorithm.SHA2_256)
                                    .setMacKey(salt.toByteString())
                                    .setMacAlgorithm(Signature.MACAlgorithm.HMAC_SHA2_256)
                                    .setSignatureLevel(Signature.SignatureLevel.ADVANCED)
                                    .setIdToken(ByteString.copyFromUtf8(input.value.id_token))
                                    .setJwkIdp(
                                        ByteString.copyFromUtf8(
                                            oidcService.marshalJwk(jwtValidationResult.jwk)
                                        )
                                    )
                                    .build().toByteArray(),
                                certificateHolder
                            ).encoded.also { pkcs7Signature ->
                                Signature.SignatureFile.newBuilder()
                                    .setSignatureDataInPkcs7(pkcs7Signature.toByteString())
                                    .addRfc3161InPkcs7(tsaService.stamp(pkcs7Signature).toByteString())
                                    .build().toByteArray()
                            }.also { signatureFile ->
                                signingKeyService.destroySigningKey(subjectInformation.value)
                                File("/tmp/signaturefile").writeBytes(signatureFile)
                                signatureHoldingService.generateId().also { id ->
                                    signatureHoldingService.set(id, signatureFile)
                                    call.respond(
                                        SigningResponse(
                                            signature =
                                            "${URLs.BASE_URL}/${locations.href(SignatureRetrievalRequest(id = id))}"
                                        )
                                    )
                                }
                            }
                        }
                    }
                    is Either.Error -> throw subjectInformation.e
                }
            }
            is Invalid -> throw input.error
        }
    }
}

