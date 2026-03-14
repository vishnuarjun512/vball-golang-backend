package ability_service

import (
	"context"
	"vball/internal/models"
	"vball/internal/repositories"
)

func CreateSubAbility(ability models.SubAbility) error {
	return repositories.CreateSubAbility(context.Background(), ability)
}

func GetSubAbilities() ([]models.SubAbility, error) {
	return repositories.GetSubAbilities(context.Background())
}

func GetSubAbility(id int) (*models.SubAbility, error) {
	return repositories.GetSubAbility(context.Background(), id)
}

func UpdateSubAbility(id int, ability models.SubAbility) error {
	return repositories.UpdateSubAbility(context.Background(), id, ability)
}

func DeleteSubAbility(id int) error {
	return repositories.DeleteSubAbility(context.Background(), id)
}
