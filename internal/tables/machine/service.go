package machine

import (
	"context"
	"vball/internal/models"
)

func GetAllMachines_Service() ([]models.Machine, error) {
	return GetAllMachines_Repo(context.Background())
}
