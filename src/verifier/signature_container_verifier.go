package verifier

import (
	"crypto/x509"
	"fmt"
	"github.com/golang/protobuf/proto"
	"go.mozilla.org/pkcs7"
)

type SignatureContainerVerifier struct {
	container *SignatureContainer
	data chan SignatureData
}

func NewSignatureContainerVerifier(c *SignatureContainer) *SignatureContainerVerifier {
	return &SignatureContainerVerifier{
		container: c,
		data:      make(chan SignatureData, 1),
	}
}

func (s SignatureContainerVerifier) Verify() error {
	p7, err :=pkcs7.Parse(s.container.EnvelopedSignatureDataPkcs7)
	if err != nil {
		return fmt.Errorf("could not decode signature container: %w", err)
	}
	signatureData := SignatureData{}
	if err := proto.Unmarshal(p7.Content, &signatureData); err != nil {
		return fmt.Errorf("could not unmarshal signature data: %w", err)
	}
	s.data <- signatureData
	certPool, err := x509.SystemCertPool()
	if err != nil {
		return fmt.Errorf("could not get system cert pool: %w", err)
	}

	if err := p7.VerifyWithChain(certPool); err != nil {
		return fmt.Errorf("could not verify pcks7 signature container: %w", err)
	}

	l := ltvVerifier{
		certs:  p7.Certificates,
		ltvMap: s.container.LtvSigning,
	}
	if err := l.Verify(); err != nil {
		return fmt.Errorf("ltv information for signature is not valid: %w", err)
	}
	return nil
}

func (s *SignatureContainerVerifier) getSignatureData() SignatureData {
	return <-s.data
}