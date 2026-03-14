package ability_service

import (
	"context"
	"vball/internal/models"
	ability_repo "vball/internal/repositories/ability"
)

func CreateMainAbility(ability models.CreateAbilityRequest) (*models.MainAbility, error) {
	createdAbility, err := ability_repo.CreateMainAbility(context.Background(), ability)
	if err != nil {
		return nil, err
	}
	return createdAbility, nil
}

func GetMainAbilities() ([]models.MainAbility, error) {
	return ability_repo.GetMainAbilities(context.Background())
}

func GetMainAbility(id int) (*models.MainAbility, error) {
	return ability_repo.GetMainAbility(context.Background(), id)
}

func UpdateMainAbility(id int, ability models.MainAbility) error {
	return ability_repo.UpdateMainAbility(context.Background(), id, ability)
}

func DeleteMainAbility(id int) error {
	return ability_repo.DeleteMainAbility(context.Background(), id)
}
