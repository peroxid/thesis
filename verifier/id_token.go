package verifier

import (
	"context"
	"errors"
	"fmt"
	"github.com/coreos/go-oidc"
	log "github.com/sirupsen/logrus"
	"gitlab.ti.bfh.ch/hirtp1/thesis/src/verifier/pb"
	"gopkg.in/square/go-jose.v2"
	"time"
)

type idTokenVerifier struct {
	token       []byte
	signerEmail chan string
	idToken     chan idToken
	notAfter    func() time.Time
	key         jose.JSONWebKey
	ltvData     map[string]*pb.LTV
	verifyLTV   bool
	ctx         context.Context
	cfg         *Config
}

type emailClaims struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

type idToken struct {
	oidc.IDToken
	emailClaims
	Certs []CertChain `json:"cert_chain"`
}

func NewIDTokenVerifier(signatureData *pb.SignatureData, notAfter time.Time, cfg Config) (*idTokenVerifier, error) {
	if signatureData == nil {
		return nil, errors.New("signature data can't be nil")
	}
	cfg.Logger = cfg.Logger.WithField("verifier", "id token")
	i := &idTokenVerifier{
		token:       signatureData.IdToken,
		signerEmail: make(chan string, 1),
		idToken:     make(chan idToken, 1),
		notAfter:    notAfter.Local,
		ltvData:     signatureData.LtvIdp,
		ctx:         context.Background(),
		key:         jose.JSONWebKey{},
		cfg:         &cfg,
	}
	if err := i.key.UnmarshalJSON(signatureData.JwkIdp); err != nil {
		return nil, fmt.Errorf("could not unmarshal jwk: %w", err)
	}
	cfg.Logger.WithFields(log.Fields{
		"not_after": notAfter,
		"issuer":    cfg.Issuer,
		"client_id": cfg.ClientId,
	}).Info("created new id token verifier")
	return i, nil
}

func (i *idTokenVerifier) VerifySignature(ctx context.Context, jwtRaw string) (payload []byte, err error) {
	signature, err := jose.ParseSigned(jwtRaw)
	if err != nil {
		return nil, fmt.Errorf("could not parse token: %w", err)
	}

	return signature.Verify(i.key)
}

func (i *idTokenVerifier) SendEmail(signerEmail string) {
	i.signerEmail <- signerEmail
}

func (i *idTokenVerifier) IDToken() idToken {
	return <-i.idToken
}

func (i *idTokenVerifier) Verify(verifyLTV bool) error {
	i.cfg.Logger.Info("started verifying")
	oidcCfg := &oidc.Config{
		ClientID: i.cfg.ClientId,
		Now:      i.notAfter,
	}
	verifier := oidc.NewVerifier(i.cfg.Issuer, i, oidcCfg)
	decodedIDToken, err := verifier.Verify(i.ctx, string(i.token))
	if err != nil {
		return err
	}
	i.cfg.Logger.WithFields(log.Fields{
		"issuer":    decodedIDToken.Issuer,
		"expiry":    decodedIDToken.Expiry,
		"issued_at": decodedIDToken.IssuedAt,
		"nonce":     decodedIDToken.Nonce,
		"audience":  decodedIDToken.Audience,
		"subject":   decodedIDToken.Subject,
	}).Info("decoded id token")

	var emailClaims emailClaims
	if err = decodedIDToken.Claims(&emailClaims); err != nil {
		return err
	}
	idTokenWithClaims := idToken{
		IDToken:     *decodedIDToken,
		emailClaims: emailClaims,
	}
	for _, c := range i.key.Certificates {
		idTokenWithClaims.Certs = append(idTokenWithClaims.Certs, CertChain{
			Issuer:    c.Issuer.String(),
			Subject:   c.Subject.String(),
			NotBefore: c.NotBefore,
			NotAfter:  c.NotAfter,
		})
	}
	i.idToken <- idTokenWithClaims

	if !emailClaims.EmailVerified {
		return errors.New("e-mail was not verified")
	}
	i.cfg.Logger.WithFields(log.Fields{
		"email":          emailClaims.Email,
		"email_verified": emailClaims.EmailVerified,
	}).Info("decoded id token claims")

	signerEmail := <-i.signerEmail
	if emailClaims.Email != signerEmail {
		return fmt.Errorf("id token email %s doesn't match signerEmail email %s", emailClaims.Email, signerEmail)
	}

	if verifyLTV {
		l := LTVVerifier{
			certs: i.key.Certificates,
			//LTVData: i.ltvData,
		}
		if err = l.Verify(); err != nil {
			return fmt.Errorf("verifyLTV information for id token not valid: %w", err)
		}
	}
	i.cfg.Logger.Info("finished verifying")

	return nil
}
