package kfk

import (
	_ "net/http/pprof"

	"github.com/IBM/sarama"
)

// KafkaProducer 发送单条
func KafkaProducer(topic string, msg []byte) (err error) {
	producer, err := sarama.NewSyncProducerFromClient(GobalKafka)
	if err != nil {
		return
	}
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}
	_, _, err = producer.SendMessage(message)
	if err != nil {
		return
	}
	return
}

// KafkaProducers 发送多条，topic在messages中
func KafkaProducers(messages []*sarama.ProducerMessage) (err error) {
	producer, err := sarama.NewSyncProducerFromClient(GobalKafka)
	if err != nil {
		return
	}
	err = producer.SendMessages(messages)
	if err != nil {
		return
	}
	return
}
