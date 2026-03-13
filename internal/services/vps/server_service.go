package vps

import (
	"fmt"
)

type ServerService struct{}

func (s *ServerService) StartServer(port int) int {

	fmt.Println("Starting Unreal server:", port)

	// simulate process id
	return port * 100
}

func (s *ServerService) StopServer(pid int) {

	fmt.Println("Stopping server process:", pid)
}
