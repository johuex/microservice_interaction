package shared

import (
	"kafka_service/config"
	service "kafka_service/services"
	"log"
	"os"
)

var ContainerItem Container

type Container struct {
	Service *service.Service
}

func init() {
	logger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	config_ := config.New()
	service_ := &service.Service{Logger: logger, Config: config_}

	ContainerItem = Container{service_}
}
