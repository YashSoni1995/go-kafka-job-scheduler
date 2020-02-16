package config

import (
	"gopkg.in/pg.v4"
	"github.com/shopify/sarama"
	"log"
)

type Config struct {
	PG     *pg.DB
	Port   string
	Env    string
	Kafka  sarama.SyncProducer
}

//initializing postgres database
func InitPG(env string) *pg.DB {
	var dbName string
	dbName = "testdb"
	log.Println("[PG] Connecting to " + "localhost" + "/" + dbName)
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "testuser",
		Password: "testuserpassword",
		Database: dbName,
	})
	return db
}

//initializing kafka producer
func InitKafkaProducer(env string) sarama.SyncProducer {
	brokers := []string{"localhost:9092"}
	var config *sarama.Config
	config = sarama.NewConfig()
	config.Producer.Return.Successes = true
	w, err := sarama.NewSyncProducer(brokers, config)
	log.Println("[KAFKA] Connecting to kafka broker", w, "with error", err)
	return w
}