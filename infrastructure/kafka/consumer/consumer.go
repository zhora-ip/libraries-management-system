package kafkaservice

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	cfg "github.com/zhora-ip/libraries-management-system/infrastructure/kafka"
)

type Consumer struct {
	consumer *kafka.Consumer
	topic    string
	done     chan struct{}
	stop     chan struct{}
}

func New(cfg *cfg.Config) (*Consumer, error) {

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":       cfg.Server,
		"log_level":               cfg.LogLevel,
		"auto.offset.reset":       cfg.OffsetReset,
		"group.id":                cfg.GroupID,
		"log.connection.close":    false,
		"enable.auto.commit":      true,
		"auto.commit.interval.ms": 5000,
	})

	if err != nil {
		return nil, err
	}

	err = consumer.Subscribe(cfg.Topic, nil)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		consumer: consumer,
		topic:    cfg.Topic,
		done:     make(chan struct{}),
		stop:     make(chan struct{}),
	}, nil
}

func (c *Consumer) GetDoneCh() <-chan struct{} {
	return c.done
}

func (c *Consumer) GetStopCh() chan<- struct{} {
	return c.stop
}

func (c *Consumer) ShutDown() {
	close(c.done)
	<-c.stop
	if err := c.consumer.Close(); err != nil {
		log.Print(err)
	}
}
