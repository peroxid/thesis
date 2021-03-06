package verifier_test

import (
	"crypto/sha256"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gitlab.ti.bfh.ch/hirtp1/thesis/src/verifier"
	"gitlab.ti.bfh.ch/hirtp1/thesis/src/verifier/pb"
	"testing"
	"time"
)

func TestVerifyIDToken(t *testing.T) {
	idTokenFile := readFile(t, "idtoken_keycloak")
	idTokenManipulatedFile := readFile(t, "idtoken_keycloak_manipulated")
	intermediateCAOCSPFile := readFile(t, "SwissSign TSA Platinum CA 2017 - G22.pem.ocsp")
	tsaCAOCSPFile := readFile(t, "SwissSign ZertES TSA UNIT CH-2018.pem.ocsp")
	intermediateCA := parsePEM(t, "SwissSign TSA Platinum CA 2017 - G22.pem")
	tsaCA := parsePEM(t, "SwissSign ZertES TSA UNIT CH-2018.pem")

	jwkFile := readFile(t, "jwk.json")
	ltv := map[string]*pb.LTV{
		fmt.Sprintf("%x", sha256.Sum256(intermediateCA.Raw)): {
			Ocsp: intermediateCAOCSPFile,
		},
		fmt.Sprintf("%x", sha256.Sum256(tsaCA.Raw)): {
			Ocsp: tsaCAOCSPFile,
		},
	}

	type args struct {
		token     []byte
		issuer    string
		nonce     string
		clientId  string
		notAfter  time.Time
		key       []byte
		ltv       map[string]*pb.LTV
		verifyLTV bool
		email     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid id token",
			args: args{
				token:     idTokenFile,
				issuer:    "https://keycloak.thesis.izolight.xyz/auth/realms/master",
				nonce:     "5093b0fb5a68144fd3fddda5156f232e975c6eb857cba5b5fd9d64b7b31bbe45",
				clientId:  "thesis",
				notAfter:  time.Unix(1575021202, 0),
				key:       jwkFile,
				ltv:       ltv,
				verifyLTV: false,
				email:     "test2@thesis.izolight.xyz",
			},
			wantErr: false,
		},
		{
			name: "valid id token without ltv",
			args: args{
				token:     idTokenFile,
				issuer:    "https://keycloak.thesis.izolight.xyz/auth/realms/master",
				nonce:     "5093b0fb5a68144fd3fddda5156f232e975c6eb857cba5b5fd9d64b7b31bbe45",
				clientId:  "thesis",
				notAfter:  time.Unix(1575021202, 0),
				key:       jwkFile,
				ltv:       ltv,
				verifyLTV: true,
				email:     "test2@thesis.izolight.xyz",
			},
			wantErr: true,
		},
		{
			name: "valid , but expired id token (okta)",
			args: args{
				token:     idTokenFile,
				issuer:    "https://keycloak.thesis.izolight.xyz/auth/realms/master",
				nonce:     "5093b0fb5a68144fd3fddda5156f232e975c6eb857cba5b5fd9d64b7b31bbe45",
				clientId:  "thesis",
				notAfter:  time.Now(),
				key:       jwkFile,
				ltv:       ltv,
				verifyLTV: true,
			},
			wantErr: true,
		},
		{
			name: "wrong nonce (okta)",
			args: args{
				token:     idTokenFile,
				issuer:    "https://keycloak.thesis.izolight.xyz/auth/realms/master",
				nonce:     "5093b0fb5a68144fd3fddda5156f232e975c6eb857cba5b5fd9d64b7b31bbea5",
				clientId:  "thesis",
				notAfter:  time.Unix(1575021202, 0),
				key:       jwkFile,
				ltv:       ltv,
				verifyLTV: true,
			},
			wantErr: true,
		},
		{
			name: "manipulated id token (okta)",
			args: args{
				token:     idTokenManipulatedFile,
				issuer:    "https://keycloak.thesis.izolight.xyz/auth/realms/master",
				nonce:     "5093b0fb5a68144fd3fddda5156f232e975c6eb857cba5b5fd9d64b7b31bbea5",
				clientId:  "thesis",
				notAfter:  time.Unix(1575021202, 0),
				key:       jwkFile,
				ltv:       ltv,
				verifyLTV: true,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := log.New()
			logger.SetLevel(log.FatalLevel)
			cfg := verifier.Config{
				Issuer:   tt.args.issuer,
				ClientId: tt.args.clientId,
				Logger:   log.NewEntry(logger),
			}
			signatureData := &pb.SignatureData{
				IdToken: tt.args.token,
				JwkIdp:  tt.args.key,
				LtvIdp:  tt.args.ltv,
			}
			v, err := verifier.NewIDTokenVerifier(signatureData, tt.args.notAfter, cfg)
			v.SendEmail(tt.args.email)
			if err != nil {
				t.Fatalf("NewIDTokenVerifier error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := v.Verify(tt.args.verifyLTV); err != nil != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}