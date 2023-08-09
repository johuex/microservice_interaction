package shared

import (
	"log"
	"master_service/config"
	"master_service/services"
	"os"
)

var ContainerItem Container

type Container struct {
	Service *services.Service
}

func init() {
	config_ := config.New()
	logger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	service := services.Service{Config: config_, Logger: logger}
	ContainerItem = Container{&service}
}
