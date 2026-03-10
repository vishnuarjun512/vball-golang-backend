package services

import (
	"context"
	"vball/internal/models"
	"vball/internal/repositories"
)

func GetAllAbilities() ([]models.MainAbility, []models.SubAbility, error) {

	return repositories.GetAllAbilities(context.Background())
}

func GetSteamLogin_Service(steamID string, username string) (*models.PlayerAdmin, error) {

	player, err := repositories.GetPlayerBySteamID_Repo(context.Background(), steamID)

	if err == nil && player != nil {
		return player, nil
	}

	playerID, err := repositories.CreatePlayer_Repo(context.Background(), steamID, username)

	if err != nil {
		return nil, err
	}

	err = repositories.CreatePlayerAbilities_Repo(context.Background(), playerID)

	if err != nil {
		return nil, err
	}

	return repositories.GetPlayerBySteamID_Repo(context.Background(), steamID)
}
