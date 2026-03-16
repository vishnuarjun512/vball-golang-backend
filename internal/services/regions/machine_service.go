package region

import (
	"context"
	"vball/internal/models"
	"vball/internal/repositories/region"
)

func GetAllMachines_Service() ([]models.Machine, error) {
	return region.GetAllMachines_Repo(context.Background())
}
