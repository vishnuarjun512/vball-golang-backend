package services

import (
	"context"
	"vball/internal/models"
	"vball/internal/repositories"
)

func GetAdminLoadOut_Service() ([]models.PlayerAdmin, error) {
	return repositories.GetAdminLoadout_Repo(context.Background())
}

func GetAllAbilities_Service() ([]models.MainAbility, []models.SubAbility, error) {
	return repositories.GetAllAbilities_Repo(context.Background())
}
