package region

import (
	"context"
	"vball/internal/models"
	"vball/internal/repositories/region"
)

func GetAllGameServers_Service() ([]models.GameServer, error) {
	return region.GetAllGameServers_Repo(context.Background())
}
