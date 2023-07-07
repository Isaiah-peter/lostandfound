package token

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Payload struct {
	ID uuid.UUID `json:"id"`
	UserId int32 `json:"user_id"`
	IssueAt time.Time `json:"issue_at"`
	ExpiriedAt time.Time `json:"expiried_at"`
}

func NewPayload(user_id int32, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID: tokenId,
		UserId: user_id,
		IssueAt: time.Now(),
		ExpiriedAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiriedAt) {
		return errors.New("payload has expired")
	}

	return nil
}