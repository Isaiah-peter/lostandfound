package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const minSecretekeySize = 32

var ErrExpiredToken = errors.New("token is expired")

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretkey string) (Maker, error) {
	if len(secretkey) < minSecretekeySize {
		return nil, fmt.Errorf("the secretkey is tool short min %v", minSecretekeySize)
	}

	return &JWTMaker{secretKey: secretkey}, nil
}

func (maker *JWTMaker) CreateToken(user_id int32, duration time.Duration) (string, error) {
	payload, err := NewPayload(user_id, duration)

	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString(maker.secretKey)
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		return []byte(maker.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)

	if err != nil {
		verr, ok := err.(*jwt.ValidationError)

		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}

		return nil, errors.New("invalid token")
	}

	payload, ok := jwtToken.Claims.(*Payload)

	if !ok {
		return nil, errors.New("invalid token")
	}

	return payload, nil
}
