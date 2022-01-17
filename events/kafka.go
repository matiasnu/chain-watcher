package events

import (
	"time"

	"github.com/matiasnu/chain-watcher/models"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

const clientID = "01DJQXRBDE49RGSZ87B126FCB0"

// NewPublisher create a kafka publisher
func NewPublisher(brokers []string, topic string) models.Publisher {
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: clientID,
	}

	c := kafka.WriterConfig{
		Brokers:          brokers,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}

	return &publisher{kafka.NewWriter(c)}
}
