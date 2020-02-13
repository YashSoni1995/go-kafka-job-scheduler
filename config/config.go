package config

import (
	"gopkg.in/pg.v4"
	"log"
)

type Config struct {
	PG     *pg.DB
	Port   string
	Env    string
	//Kafka  sarama.SyncProducer
}

func InitPG(env string) *pg.DB {
	var dbName string
	dbName = "test"
	log.Println("[PG] Connecting to " + "localhost" + "/" + dbName)
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "testuser",
		Password: "testpassword",
		Database: dbName,
	})
	return db
}