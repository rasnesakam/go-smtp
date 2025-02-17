package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

type EventHandler func([]byte)

func readEvents(eventHandler EventHandler) error {
	fmt.Println("Setting up Kafka Environments")
	err := godotenv.Load()
	if err != nil {
		return errors.New("error loading .env file")
	}
	kafkaBrokers, ok := os.LookupEnv("KAFKA_BROKERS")
	if !ok {
		return errors.New("undefined environment KAFKA_BROKERS")
	}
	kafkaTopic, ok := os.LookupEnv("KAFKA_TOPIC")
	if !ok {
		return errors.New("undefined environment KAFKA_TOPIC")
	}
	kafkaPartition, ok := os.LookupEnv("KAFKA_PARTITION")
	if !ok {
		return errors.New("undefined environment KAFKA_PARTITION")
	}
	kafkaGroupID, ok := os.LookupEnv("KAFKA_GROUP_ID")
	if !ok {
		return errors.New("undefined environment KAFKA_GROUP_ID")
	}
	partition, err := strconv.Atoi(kafkaPartition)
	if err != nil {
		return errors.New("KAFKA_PARTITION value should be an integer")
	}
	fmt.Println("Connecting to kafka")

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   strings.Split(kafkaBrokers, ","),
		Topic:     kafkaTopic,
		Partition: partition,
		MaxBytes:  10e6,
		GroupID:   kafkaGroupID,
	})
	ctx := context.Background()
	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			break
		}
		eventHandler(m.Value)
		if err := r.CommitMessages(ctx, m); err != nil {
			fmt.Println(err)
			break
		}
	}

	if err := r.Close(); err != nil {
		return err
	}

	return err
}
