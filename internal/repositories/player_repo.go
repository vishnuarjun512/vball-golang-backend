package repositories

import (
	"context"
	"fmt"
	"vball/internal/database"
	"vball/internal/models"
)

func CreatePlayer_Repo(ctx context.Context, steamID string, username string) (string, error) {

	query := `
	INSERT INTO players (steam_id, username)
	VALUES ($1, $2)
	RETURNING player_id
	`

	var playerID string

	err := database.DB.QueryRow(ctx, query, steamID, username).Scan(&playerID)

	if err != nil {
		fmt.Printf("Error creating player: %v\n", err)
		return "", err
	}

	return playerID, nil
}

func CreatePlayerAbilities_Repo(ctx context.Context, playerID string) error {

	query := `
	INSERT INTO player_abilities
	(player_id)
	VALUES ($1)
	`

	_, err := database.DB.Exec(ctx, query, playerID)

	if err != nil {
		fmt.Printf("Error creating player abilities: %v\n", err)
	}

	return err
}

func GetAdminLoadout_Repo(ctx context.Context) ([]models.PlayerAdmin, error) {

	query := `
	SELECT
	p.player_id,
	p.steam_id,
	p.username,
	p.kash,
	p.ban_status,
	p.matches_played,
	p.wins,
	p.losses,
	p.last_login,
	p.created_at,

	pa.main_ability_id,
	pa.sub_ability_slot1,
	pa.sub_ability_slot2,
	pa.sub_ability_slot3

	FROM players p
	LEFT JOIN player_abilities pa
	ON p.player_id = pa.player_id
	`

	rows, err := database.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var players []models.PlayerAdmin

	for rows.Next() {

		var p models.PlayerAdmin
		var sub1, sub2, sub3 *int

		err := rows.Scan(
			&p.PlayerID,
			&p.SteamID,
			&p.Username,
			&p.Kash,
			&p.BanStatus,
			&p.MatchesPlayed,
			&p.Wins,
			&p.Losses,
			&p.LastLogin,
			&p.RegisteredAt,
			&p.MainAbilityID,
			&sub1,
			&sub2,
			&sub3,
		)

		if err != nil {
			return nil, err
		}

		var subAbilities []int

		if sub1 != nil {
			subAbilities = append(subAbilities, *sub1)
		}
		if sub2 != nil {
			subAbilities = append(subAbilities, *sub2)
		}
		if sub3 != nil {
			subAbilities = append(subAbilities, *sub3)
		}

		p.SubAbilityIDs = subAbilities

		players = append(players, p)
	}

	return players, nil
}

func GetPlayerBySteamID_Repo(ctx context.Context, steamID string) (*models.PlayerAdmin, error) {

	query := `
	SELECT
	p.player_id,
	p.steam_id,
	p.username,
	p.kash,
	p.ban_status,
	p.matches_played,
	p.wins,
	p.losses,
	p.last_login,
	p.created_at,

	pa.main_ability_id,
	pa.sub_ability_slot1,
	pa.sub_ability_slot2,
	pa.sub_ability_slot3

	FROM players p

	LEFT JOIN player_abilities pa
	ON p.player_id = pa.player_id

	WHERE p.steam_id = $1
	`

	row := database.DB.QueryRow(ctx, query, steamID)

	var player models.PlayerAdmin
	var sub1, sub2, sub3 *int

	err := row.Scan(
		&player.PlayerID,
		&player.SteamID,
		&player.Username,
		&player.Kash,
		&player.BanStatus,
		&player.MatchesPlayed,
		&player.Wins,
		&player.Losses,
		&player.LastLogin,
		&player.RegisteredAt,
		&player.MainAbilityID,
		&sub1,
		&sub2,
		&sub3,
	)

	if err != nil {
		return nil, err
	}

	var subAbilities []int

	if sub1 != nil {
		subAbilities = append(subAbilities, *sub1)
	}
	if sub2 != nil {
		subAbilities = append(subAbilities, *sub2)
	}
	if sub3 != nil {
		subAbilities = append(subAbilities, *sub3)
	}

	player.SubAbilityIDs = subAbilities

	return &player, nil
}
