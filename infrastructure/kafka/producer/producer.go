package kafkaservice

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	cfg "github.com/zhora-ip/libraries-management-system/infrastructure/kafka"
)

const (
	flushTimeout  = 1000
	lenDeliveryCh = 1000
)

type Producer struct {
	producer   *kafka.Producer
	topic      string
	deliveryCh chan kafka.Event
}

func New(cfg *cfg.Config) (*Producer, error) {

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":    cfg.Server,
		"log_level":            cfg.LogLevel,
		"client.id":            "log-client",
		"delivery.timeout.ms":  5000,
		"log.connection.close": false,
		"acks":                 "all",
	})

	if err != nil {
		return nil, err
	}

	return &Producer{
		producer:   p,
		topic:      cfg.Topic,
		deliveryCh: make(chan kafka.Event, lenDeliveryCh),
	}, nil
}

func (p *Producer) ShutDown() {
	p.producer.Flush(flushTimeout)
	p.producer.Close()
}
