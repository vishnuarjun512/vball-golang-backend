package player

import (
	"context"
	"database/sql"
	"errors"
	"vball/internal/models"
)

func GetSteamLogin_Service(steamID string, username string) (*models.PlayerAdmin, error) {
	ctx := context.Background()

	// 1. Try to find the player
	player, err := GetPlayerBySteamID_Repo(ctx, steamID)

	// 2. If no error, return the player
	if err == nil {
		return player, nil
	}

	// 3. If error is NOT "sql.ErrNoRows", something went wrong with DB
	if !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	// 4. Player doesn't exist -> Create them
	playerID, err := CreatePlayer_Repo(ctx, steamID, username)
	if err != nil {
		return nil, err
	}

	// 5. Initialize their abilities table
	if err := CreatePlayerAbilities_Repo(ctx, playerID); err != nil {
		return nil, err
	}

	// 6. Return the newly created player
	return GetPlayerBySteamID_Repo(ctx, steamID)
}

func GetPlayerBySteamID_Service(steamID string) (*models.PlayerAdmin, error) {
	return GetPlayerBySteamID_Repo(context.Background(), steamID)
}

func CreatePlayer_Service(steamID string, username string) (string, error) {
	ctx := context.Background()
	return CreatePlayer_Repo(ctx, steamID, username)
}

func DeletePlayer_Service(steamId string) error {
	ctx := context.Background()
	return DeletePlayer_Repo(ctx, steamId)
}
