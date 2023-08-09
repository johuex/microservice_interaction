package service

import (
	"context"
	"log"
	"sync"

	"kafka_service/config"

	"github.com/segmentio/kafka-go"
)

type Service struct {
	Logger *log.Logger
	Config *config.Config
}

func (s *Service) Consume(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{s.Config.BrokerUrl + ":" + s.Config.BrokerPort},
		Topic:   s.Config.BrokerTopic,
	})
	ctx := context.Background()
	s.Logger.Print("Connected to Kafka as Consumer")
	for {

		msg, err := r.ReadMessage(ctx) // block & wait for NewMessage
		if err != nil {
			//wg.Done()
			s.Logger.Fatal("could not read message " + err.Error())
		}
		s.Logger.Printf("Received from master %s:%s ", string(msg.Key), string(msg.Value))
	}
}
