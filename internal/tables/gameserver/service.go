package gameserver

import (
	"context"
	"fmt"
	"vball/internal/models"
)

func GetAllGameServers_Service() ([]models.GameServer, error) {
	return GetAllGameServers_Repo(context.Background())
}

func CreateGameServer_Service(machineId int, port int, maxPlayers int) (*models.GameServer, error) {
	return CreateGameServer_Repo(context.Background(), machineId, port, maxPlayers)
}

func GetGameServer_Service(id int) (*models.GameServer, error) {
	return GetGameServer_Repo(context.Background(), id)
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
