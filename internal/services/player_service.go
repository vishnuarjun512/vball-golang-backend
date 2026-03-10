package services

import (
	"context"
	"vball/internal/models"
	"vball/internal/repositories"
)

func GetAllPlayersLoadOut_Service() ([]models.PlayerAdmin, error) {

	return repositories.GetAllPlayersLoadOut_Repo(context.Background())
}

func GetPlayerBySteamID_Service(steamID string) (*models.PlayerAdmin, error) {

	return repositories.GetPlayerBySteamID_Repo(context.Background(), steamID)
}
