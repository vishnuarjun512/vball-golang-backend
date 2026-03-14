package ability_service

import (
	"context"
	"vball/internal/models"
	ability_repo "vball/internal/repositories/ability"
)

func CreateSubAbility(ability models.SubAbility) error {
	return ability_repo.CreateSubAbility(context.Background(), ability)
}

func GetSubAbilities() ([]models.SubAbility, error) {
	return ability_repo.GetSubAbilities(context.Background())
}

func GetSubAbility(id int) (*models.SubAbility, error) {
	return ability_repo.GetSubAbility(context.Background(), id)
}

func UpdateSubAbility(id int, ability models.SubAbility) error {
	return ability_repo.UpdateSubAbility(context.Background(), id, ability)
}

func DeleteSubAbility(id int) error {
	return ability_repo.DeleteSubAbility(context.Background(), id)
}
