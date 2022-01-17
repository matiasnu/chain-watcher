package events

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/segmentio/kafka-go"
)

type publisher struct {
	writer *kafka.Writer
}

func (p *publisher) Publish(ctx context.Context, payload interface{}) error {
	message, err := p.encodeMessage(payload)
	if err != nil {
		return err
	}

	return p.writer.WriteMessages(ctx, message)
}

func (p *publisher) encodeMessage(payload interface{}) (kafka.Message, error) {
	m, err := json.Marshal(payload)
	if err != nil {
		return kafka.Message{}, err
	}

	key := Ulid()
	return kafka.Message{
		Key:   []byte(key),
		Value: m,
	}, nil
}

// Ulid encapsulate the way to generate ulids
func Ulid() string {
	t := time.Now().UTC()
	id := ulid.MustNew(ulid.Timestamp(t), rand.Reader)

	return id.String()
}
