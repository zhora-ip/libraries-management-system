package kafkaservice

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

func (c *Consumer) Poll(timeoutMs int) ([]byte, error) {

	ev := c.consumer.Poll(timeoutMs)
	switch e := ev.(type) {
	case *kafka.Message:
		return e.Value, nil
	case *kafka.Error:
		return nil, e
	default:
		return nil, nil
	}
}
