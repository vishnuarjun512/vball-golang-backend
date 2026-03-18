package repositories

import (
	"context"
	"vball/internal/database"
	"vball/internal/models"

	"github.com/jackc/pgx/v5"
)

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

func GetAllAbilities_Repo(ctx context.Context) ([]models.MainAbility, []models.SubAbility, error) {

	mainRows, err := database.DB.Query(ctx, `
		SELECT
			id,
			name,
			description,
			type,
			tier,
			duration,
			cooldown,
			spike_modifier,
			jump_modifier,
			set_modifier,
			receive_modifier,
			ball_force_multiplier
		FROM main_abilities
	`)

	if err != nil {
		return nil, nil, err
	}

	defer mainRows.Close()

	mainAbilities, err := pgx.CollectRows(mainRows, pgx.RowToStructByName[models.MainAbility])

	if err != nil {
		return nil, nil, err
	}

	subRows, err := database.DB.Query(ctx, `
		SELECT
			id,
			name,
			description,
			tier,
			modifier_type,
			modifier_value
		FROM sub_abilities
	`)

	if err != nil {
		return nil, nil, err
	}

	defer subRows.Close()

	subAbilities, err := pgx.CollectRows(subRows, pgx.RowToStructByName[models.SubAbility])

	if err != nil {
		return nil, nil, err
	}

	return mainAbilities, subAbilities, nil
}
