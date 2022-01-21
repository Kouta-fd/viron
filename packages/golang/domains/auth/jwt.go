package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cam-inc/viron/packages/golang/logging"

	"github.com/cam-inc/viron/packages/golang/errors"

	"github.com/cam-inc/viron/packages/golang/constant"

	"github.com/lestrrat-go/jwx/jwa"

	"github.com/go-chi/jwtauth"
)

type (
	JWT struct {
		Secret        string
		Provider      func(r *http.Request) (string ,[]string, error)
		ExpirationSec int
		jwtAuth       *jwtauth.JWTAuth
	}
	Config struct {
		Secret        string
		Provider      string
		ExpirationSec int
	}
	Claim struct {
		Exp int
		Iat int
		Nbf int
		Sub string
		Iss string
		Aud []string
	}
)

var (
	jwt *JWT
	log logging.Logger
)

func SetUp(secret string, provider func(r *http.Request) (string ,[]string, error), expiration int) error {
	jwt = &JWT{
		Secret:        secret,
		Provider:      provider,
		ExpirationSec: expiration,
		jwtAuth:       jwtauth.New(string(jwa.HS512), []byte(secret), nil),
	}
	log = logging.GetDefaultLogger()
	return nil
}

func Sign(r *http.Request, subject string) (string, error) {
	iss, aud, err := jwt.Provider(r)
	if err != nil {
		return "", err
	}
	claim := map[string]interface{}{
		"nbf": 0,
		"sub": subject,
		"iss": iss,
		"aud": aud,
	}
	jwtauth.SetExpiryIn(claim, time.Duration(jwt.ExpirationSec)*time.Second)
	jwtauth.SetIssuedNow(claim)
	_, tokenStr, err := jwt.jwtAuth.Encode(claim)

	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s", constant.AUTH_SCHEME, tokenStr), nil
}

func Verify(token string) (*Claim, error) {

	if jwt == nil {
		return nil, errors.JwtUninitialized
	}

	if IsSignedOut(context.Background(), token) {
		return nil, fmt.Errorf("this token is revoked %s", token)
	}

	jwtToken, err := jwtauth.VerifyToken(jwt.jwtAuth, token)
	if err != nil {
		return nil, err
	}

	claim := &Claim{
		Exp: int(jwtToken.Expiration().Unix()),
		Iat: int(jwtToken.IssuedAt().Unix()),
		Nbf: int(jwtToken.NotBefore().Unix()),
		Sub: jwtToken.Subject(),
		Iss: jwtToken.Issuer(),
		Aud: jwtToken.Audience(),
	}

	return claim, nil
}
