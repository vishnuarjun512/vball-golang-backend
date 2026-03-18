package player

import (
	"context"
	"vball/internal/database"
	"vball/internal/models"
)

func CreatePlayer_Repo(ctx context.Context, steamID string, username string) (string, error) {
	query := `INSERT INTO players (steam_id, username) VALUES ($1, $2) RETURNING player_id`
	var playerID string
	err := database.DB.QueryRow(ctx, query, steamID, username).Scan(&playerID)
	return playerID, err
}

func CreatePlayerAbilities_Repo(ctx context.Context, playerID string) error {
	query := `INSERT INTO player_abilities (player_id) VALUES ($1)`
	_, err := database.DB.Exec(ctx, query, playerID)
	return err
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

func DeletePlayer_Repo(ctx context.Context, steamId string) error {
	query := `DELETE FROM players WHERE steamId=$1`
	_, err := database.DB.Exec(ctx, query, steamId)
	return err
}
