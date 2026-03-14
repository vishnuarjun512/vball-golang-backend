package repositories

import (
	"context"
	"vball/internal/database"
	"vball/internal/models"

	"github.com/jackc/pgx/v5"
)

func GetAllAbilities(ctx context.Context) ([]models.MainAbility, []models.SubAbility, error) {

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
