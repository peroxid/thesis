package verifier_test

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	log "github.com/sirupsen/logrus"
	"gitlab.ti.bfh.ch/hirtp1/thesis/src/verifier"
	"gitlab.ti.bfh.ch/hirtp1/thesis/src/verifier/pb"
	"testing"
)

func TestVerifySignatureData(t *testing.T) {
	type args struct {
		data      *pb.SignatureData
		hash      string
		nonce     string
		verifyLTV bool
	}
	tests := []struct {
		name string
		args
		wantErr bool
	}{
		{
			name: "valid signature with one document",
			args: args{
				hash: "87428fc522803d31065e7bce3cf03fe475096631e5e07bbd7a0fde60c4cf25c7",
			},
			wantErr: false,
		},
	}
	logger := log.New()
	logger.SetLevel(log.FatalLevel)
	cfg := verifier.Config{
		Logger: log.NewEntry(logger),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.data == nil {
				tt.args.data = generateFakeSignatureData(t, tt.hash)
			}
			v := verifier.NewSignatureDataVerifier(tt.args.data, tt.args.hash, cfg)
			if tt.nonce == "" {
				tt.nonce = generateNonce(t, tt.args.data.SaltedDocumentHash)
			}
			v.SendNonce(tt.args.nonce)
			if err := v.Verify(); err != nil != tt.wantErr {
				t.Errorf("verifiy signature data error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func generateMac(t *testing.T, hash string, key []byte) []byte {
	t.Helper()

	hashBytes, err := hex.DecodeString(hash)
	if err != nil {
		t.Fatal(err)
	}

	h := hmac.New(sha256.New, key)
	h.Write(hashBytes)
	return h.Sum(nil)
}

func generateNonce(t *testing.T, saltedHashes [][]byte) string {
	t.Helper()
	hasher := sha256.New()
	for _, m := range saltedHashes {
		hasher.Write(m)
	}
	return hex.EncodeToString(hasher.Sum(nil))
}

func generateFakeSignatureData(t *testing.T, hash string) *pb.SignatureData {
	macKey, _ := hex.DecodeString("68267cf6c2869a826d89867fd280bcdd47b33c66ef9695aac1a92e7d2a111c80")
	return &pb.SignatureData{
		HashAlgorithm:      pb.HashAlgorithm_SHA2_256,
		MacKey:             macKey,
		MacAlgorithm:       pb.MACAlgorithm_HMAC_SHA2_256,
		SaltedDocumentHash: [][]byte{generateMac(t, hash, macKey)},
	}
}
