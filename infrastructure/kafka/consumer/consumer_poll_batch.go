package kafkaservice

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	maxBatchSize = 100
)

func (c *Consumer) PollBatch(timeoutMs int) ([][]byte, error) {
	var messages [][]byte
	startTime := time.Now()
	maxTime := time.Duration(timeoutMs) * time.Millisecond

	for len(messages) < maxBatchSize {
		if time.Since(startTime) > maxTime {
			break
		}

		remainingTime := max(maxTime-time.Since(startTime), 0)

		ev := c.consumer.Poll(int(remainingTime.Milliseconds()))
		if ev == nil {
			continue
		}

		switch e := ev.(type) {
		case *kafka.Message:
			messages = append(messages, e.Value)
		case *kafka.Error:
			return messages, e
		}

		if len(messages) > 0 && time.Since(startTime) > (maxTime/2) {
			break
		}
	}

	return messages, nil
}
