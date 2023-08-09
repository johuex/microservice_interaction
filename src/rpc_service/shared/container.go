package shared

import (
	"log"
	"os"
	"rpc_service/services"
)

var ContainerItem Container

type Container struct {
	Service *services.RPCService
}

func init() {
	logger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	service := services.RPCService{Logger: logger}
	ContainerItem = Container{&service}
}
