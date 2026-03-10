package services

import (
	"context"
	"fmt"
	"vball/internal/models"
	"vball/internal/repositories"
)

func GetAllAbilities() ([]models.MainAbility, []models.SubAbility, error) {

	return repositories.GetAllAbilities(context.Background())
}

func GetSteamLogin_Service(steamID string, username string) (*models.PlayerAdmin, error) {

	player, err := repositories.GetPlayerBySteamID_Repo(context.Background(), steamID)

	if err == nil && player != nil {
		fmt.Printf("Player with SteamID %s already exists. Returning existing player.\n", steamID)
		return player, nil
	}

	playerID, err := repositories.CreatePlayer_Repo(context.Background(), steamID, username)

	if err != nil {
		fmt.Printf("Error creating player: %v\n", err)
		return nil, err
	}

	err = repositories.CreatePlayerAbilities_Repo(context.Background(), playerID)

	if err != nil {
		fmt.Printf("Error creating player abilities: %v\n", err)
		return nil, err
	}

	return repositories.GetPlayerBySteamID_Repo(context.Background(), steamID)
}
