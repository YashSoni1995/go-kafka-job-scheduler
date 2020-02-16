package kafka

import (
	"go-kafka-job-scheduler/config"
	"github.com/shopify/sarama"
	"log"
)

func PushMessageToKafka(config config.Config, topic string, message string) error {
	kafkaProducer := config.Kafka
	msg := &sarama.ProducerMessage{
		Topic: "test",
		Value: sarama.StringEncoder(message),
	}
	_, _, err := kafkaProducer.SendMessage(msg)
	if err != nil {
		log.Println("Error: %s in producing msg to %s", err, topic)
	}
	return err
}