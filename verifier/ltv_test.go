package verifier_test

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"gitlab.ti.bfh.ch/hirtp1/thesis/src/verifier"
	"testing"
)

func TestVerifyLTV(t *testing.T) {
	rootCA := parsePEM(t, "SwissSign Platinum CA - G2.pem")
	intermediateCA := parsePEM(t, "SwissSign TSA Platinum CA 2017 - G22.pem")

	intermediateCAOCSPFile := readFile(t, "SwissSign TSA Platinum CA 2017 - G22.pem.ocsp")
	tsaCAOCSPFile := readFile(t, "SwissSign ZertES TSA UNIT CH-2018.pem.ocsp")

	silverCA := parsePEM(t, "SwissSign Silver CA - G2.pem")
	revokedIntermediateCA := parsePEM(t, "SwissSign Personal Silver CA 2014 - G22.pem")
	revokedIntermediateOCSPFile := readFile(t, "SwissSign Personal Silver CA 2014 - G22.pem.ocsp")

	type args struct {
		certs []*x509.Certificate
		crls []pkix.CertificateList
		ocspResponses [][]byte
	}
	tests := []struct {
		name     string
		args args
		wantErr  bool
	}{
		{
			name: "root CA",
			args:args{
				certs:[]*x509.Certificate{rootCA},
			},
			wantErr: false,
		},
		{
			name: "intermediate CA without verifyLTV info",
			args: args{
				certs:   []*x509.Certificate{rootCA, intermediateCA},
			},
			wantErr: true,
		},
		{
			name: "intermediate CA with ocsp response",
			args: args{
				certs: []*x509.Certificate{rootCA, intermediateCA},
				ocspResponses: [][]byte{intermediateCAOCSPFile},
			},
			wantErr: false,
		},
		{
			name: "intermediate CA with ocsp response and different ca order",
			args: args{
				certs: []*x509.Certificate{intermediateCA, rootCA},
				ocspResponses: [][]byte{intermediateCAOCSPFile},
			},
			wantErr: false,
		},
		{
			name: "intermediate CA with wrong ocsp response",
			args: args{
				certs: []*x509.Certificate{rootCA, intermediateCA},
				ocspResponses: [][]byte{tsaCAOCSPFile},
			},
			wantErr: true,
		},
		{
			name: "revoked ca",
			args: args{
				certs: []*x509.Certificate{silverCA, revokedIntermediateCA},
				ocspResponses: [][]byte{revokedIntermediateOCSPFile},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := verifier.NewLTVVerifier(tt.args.certs, tt.args.crls, tt.args.ocspResponses)
			if err := v.Verify(); err != nil != tt.wantErr {
				t.Errorf("VerifyLTV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func parsePEM(t *testing.T, filename string) *x509.Certificate {
	t.Helper()
	file := readFile(t, filename)
	filePEM, _ := pem.Decode(file)
	cert, err := x509.ParseCertificate(filePEM.Bytes)
	if err != nil {
		t.Errorf("could not parse pem: %s", err)
	}
	return cert
}
