package main

import (
	"kafka_service/shared"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go shared.ContainerItem.Service.Consume(&wg)
	wg.Wait()
}
