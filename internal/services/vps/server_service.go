package vps

import (
	"context"
	"fmt"
	"vball/internal/models"
	"vball/internal/repositories/vps"
)

func CreateGameServer_Service(machineId int, port int, maxPlayers int) (*models.GameServer, error) {
	return vps.CreateGameServer_Repo(context.Background(), machineId, port, maxPlayers)
}

func GetGameServer_Service(id int) (*models.GameServer, error) {
	return vps.GetGameServer_Repo(context.Background(), id)
}

type ServerService struct{}

func (s *ServerService) StartServer(port int) int {

	fmt.Println("Starting Unreal server:", port)

	// simulate process id
	return port * 100
}

func (s *ServerService) StopServer(pid int) {

	fmt.Println("Stopping server process:", pid)
}
