package player

import (
	"context"
	"fmt"
	"vball/internal/models"
)

func GetSteamLogin_Service(steamID string, username string) (*models.PlayerAdmin, error) {

	player, err := GetPlayerBySteamID_Repo(context.Background(), steamID)

	if err == nil && player != nil {
		fmt.Printf("Player with SteamID %s already exists. Returning existing player.\n", steamID)
		return player, nil
	}

	playerID, err := CreatePlayer_Repo(context.Background(), steamID, username)

	if err != nil {
		fmt.Printf("Error creating player: %v\n", err)
		return nil, err
	}

	err = CreatePlayerAbilities_Repo(context.Background(), playerID)

	if err != nil {
		fmt.Printf("Error creating player abilities: %v\n", err)
		return nil, err
	}

	return GetPlayerBySteamID_Repo(context.Background(), steamID)
}

func GetPlayerBySteamID_Service(steamID string) (*models.PlayerAdmin, error) {
	return GetPlayerBySteamID_Repo(context.Background(), steamID)
}
