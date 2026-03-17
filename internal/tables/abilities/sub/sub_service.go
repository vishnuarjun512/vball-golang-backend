package subAbility

import (
	"context"
	"vball/internal/models"
)

func CreateSubAbility_Service(ability models.SubAbility) error {
	return CreateSubAbility_Repo(context.Background(), ability)
}

func GetSubAbilities_Service() ([]models.SubAbility, error) {
	return GetSubAbilities_Repo(context.Background())
}

func GetSubAbility_Service(id int) (*models.SubAbility, error) {
	return GetSubAbility_Repo(context.Background(), id)
}

func UpdateSubAbility_Service(id int, ability models.SubAbility) error {
	return UpdateSubAbility_Repo(context.Background(), id, ability)
}

func DeleteSubAbility_Service(id int) error {
	return DeleteSubAbility_Repo(context.Background(), id)
}
