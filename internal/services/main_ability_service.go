package services

import (
	"context"
	"vball/internal/models"
	"vball/internal/repositories"
)

func CreateMainAbility(ability models.CreateAbilityRequest) (*models.MainAbility, error) {
	createdAbility, err := repositories.CreateMainAbility(context.Background(), ability)
	if err != nil {
		return nil, err
	}
	return createdAbility, nil
}

func GetMainAbilities() ([]models.MainAbility, error) {
	return repositories.GetMainAbilities(context.Background())
}

func GetMainAbility(id int) (*models.MainAbility, error) {
	return repositories.GetMainAbility(context.Background(), id)
}

func UpdateMainAbility(id int, ability models.MainAbility) error {
	return repositories.UpdateMainAbility(context.Background(), id, ability)
}

func DeleteMainAbility(id int) error {
	return repositories.DeleteMainAbility(context.Background(), id)
}
