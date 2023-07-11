package token

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/chacha20poly1305"

	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto      *paseto.V2
	symetrickey []byte
}

var ErrInvalidToken = errors.New("token is invalid")

func NewPasetomaker(symetrickey string) (Maker, error) {
	if len(symetrickey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid keysize: must have this number of character: %d", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:      paseto.NewV2(),
		symetrickey: []byte(symetrickey),
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(user_id int32, duration time.Duration) (string, error) {
	payload, err := NewPayload(user_id, duration)

	if err != nil {
		return "", err
	}

	return maker.paseto.Encrypt(maker.symetrickey, payload, nil)
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symetrickey, payload, nil)

	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()

	if err != nil {
		return nil, err
	}

	return payload, nil
}
