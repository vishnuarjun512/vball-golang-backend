package mainAbility

import (
	"context"
)

func CreateMainAbility_Service(ability CreateAbilityRequest) (*MainAbility, error) {
	createdAbility, err := CreateMainAbility_Repo(context.Background(), ability)
	if err != nil {
		return nil, err
	}
	return createdAbility, nil
}

func GetMainAbilities_Service() ([]MainAbility, error) {
	return GetMainAbilities_Repo(context.Background())
}

func GetMainAbility_Service(id int) (*MainAbility, error) {
	return GetMainAbility_Repo(context.Background(), id)
}

func UpdateMainAbility_Service(id int, ability MainAbility) error {
	return UpdateMainAbility_Repo(context.Background(), id, ability)
}

func DeleteMainAbility_Service(id int) error {
	return DeleteMainAbility_Repo(context.Background(), id)
}
