package services

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"master_service/config"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	pb "rpc_compiled"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	Config *config.Config
	Logger *log.Logger
}

func (s *Service) Ping() []byte {
	res, err := json.Marshal(s.Config)
	if err != nil {
		s.Logger.Fatalf("Config encoding error: %v", err)
	}
	return res
}

func (s *Service) RPC() {
	// RPC call
	conn, err := grpc.Dial(s.Config.RPCServiceUrl+":"+s.Config.RPCPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		s.Logger.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	cl := pb.NewExampleServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	randInt := rand.Intn(255)
	s.Logger.Printf("RandomNumber: %d", randInt)
	r, err := cl.Call(ctx, &pb.Request{Message: int32(randInt)})
	if err != nil {
		s.Logger.Fatalf("RPC error: %v", err)
	}
	s.Logger.Printf("RPC return: %d", r.GetMessage())
}

func (s *Service) Kafka() {
	// send to service through KAFKA
	w := &kafka.Writer{
		Addr:     kafka.TCP(s.Config.KafkaURL + ":" + s.Config.KafkaPort),
		Topic:    s.Config.KafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}

	messageKey := uuid.New().String()
	randInt := rand.Intn(255)

	messageVal := strconv.Itoa(randInt)
	err := w.WriteMessages(
		context.Background(),
		kafka.Message{
			Key:   []byte(messageKey),
			Value: []byte(messageVal)},
	)
	if err != nil {
		s.Logger.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		s.Logger.Fatal("failed to close writer:", err)
	}
	s.Logger.Printf("Send to Kafka: %s:%s", messageKey, messageVal)
}

func (s *Service) API() {
	// HTTP API call
	dict := map[string]int{"randomNumber": rand.Intn(255)}
	s.Logger.Printf("RandomNumber: %d", dict["randomNumber"])
	postBody, _ := json.Marshal(dict)
	responseBody := bytes.NewBuffer(postBody)
	res, err := http.Post("http://"+s.Config.ApiServiceUrl+":"+s.Config.ApiPort, "application/json", responseBody)
	if err != nil {
		s.Logger.Fatalln(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		s.Logger.Fatalln(err)
	}
	sb := string(body)
	s.Logger.Print(sb)
}
