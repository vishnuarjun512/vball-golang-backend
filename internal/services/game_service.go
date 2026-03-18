package services

import (
	"context"
	"vball/internal/models"
	"vball/internal/repositories"

	mainAbility "vball/internal/tables/abilities/main"
	subAbility "vball/internal/tables/abilities/sub"
)

func GetAdminLoadOut_Service() ([]models.PlayerAdmin, error) {
	return repositories.GetAdminLoadout_Repo(context.Background())
}

func GetAllAbilities_Service() ([]mainAbility.MainAbility, []subAbility.SubAbility, error) {
	return repositories.GetAllAbilities_Repo(context.Background())
}
