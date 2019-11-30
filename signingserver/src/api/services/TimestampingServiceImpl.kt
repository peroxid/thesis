package ch.bfh.ti.hirtp1ganzg1.thesis.api.services

import ch.bfh.ti.hirtp1ganzg1.thesis.api.utils.SHA256
import ch.bfh.ti.hirtp1ganzg1.thesis.api.utils.byteArrayToHexString
import io.ktor.client.HttpClient
import io.ktor.client.features.logging.LogLevel
import io.ktor.client.features.logging.Logging
import io.ktor.client.request.post
import io.ktor.client.request.url
import io.ktor.http.ContentType
import io.ktor.http.content.ByteArrayContent
import org.bouncycastle.cms.CMSAlgorithm
import org.bouncycastle.tsp.TimeStampRequestGenerator
import org.slf4j.LoggerFactory
import java.security.MessageDigest

class TimestampingServiceImpl : ITimestampingService {
    companion object {
        const val TSA_URL = "http://tsa.swisssign.net"
    }

    private val logger = LoggerFactory.getLogger(this.javaClass)

    override suspend fun stamp(dataToStamp: ByteArray): ByteArray {
        val encodedTimestampingRequest = TimeStampRequestGenerator().generate(
            CMSAlgorithm.SHA256,
            MessageDigest.getInstance(SHA256).digest(dataToStamp)
        ).encoded
        val response = HttpClient {
            install(Logging) {
                level = LogLevel.HEADERS
            }
        }.use {
            it.post<ByteArray> {
                url(TSA_URL)
                body = ByteArrayContent(
                    encodedTimestampingRequest,
                    contentType = ContentType("application", "timestamp-query")
                )
            }
        }
        logger.debug("Timestamp response: {}", byteArrayToHexString(response))
        return response

    }
}