package kafkaservice

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

func (p *Producer) Produce(value []byte) error {
	err := p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &p.topic,
			Partition: kafka.PartitionAny,
		},
		Value: value,
	},
		p.deliveryCh,
	)

	if err != nil {
		return err
	}

	m := (<-p.deliveryCh).(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}

	return nil
}
