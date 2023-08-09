package services

import "log"

type RPCService struct {
	Logger *log.Logger
}

func (s *RPCService) Call(number int) int {
	s.Logger.Print(number)
	number += 1
	s.Logger.Print(number)
	return number
}
