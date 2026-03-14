package repositories

import (
	"context"
	"fmt"
	"vball/internal/database"
	"vball/internal/models"

	"github.com/jackc/pgx/v5"
)

func CreateMainAbility(ctx context.Context, ability models.CreateAbilityRequest) (*models.MainAbility, error) {

	query := `
	INSERT INTO main_abilities
	(name, description, type, tier, duration, cooldown,
	spike_modifier, jump_modifier, set_modifier, receive_modifier, ball_force_multiplier)

	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
	RETURNING id
	`

	var id int
	err := database.DB.QueryRow(ctx, query,
		ability.Name,
		ability.Description,
		ability.Type,
		ability.Tier,
		ability.Duration,
		ability.Cooldown,
		ability.SpikeModifier,
		ability.JumpModifier,
		ability.SetModifier,
		ability.ReceiveModifier,
		ability.BallForceMultiplier,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	// Fetch the created ability to return it
	return GetMainAbility(ctx, id)
}

func GetMainAbilities(ctx context.Context) ([]models.MainAbility, error) {

	query := `
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
	`

	rows, err := database.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	abilities, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MainAbility])

	if err != nil {
		return nil, err
	}

	return abilities, nil
}

func GetMainAbility(ctx context.Context, id int) (*models.MainAbility, error) {

	query := `
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
	WHERE id=$1
	`

	rows, err := database.DB.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	abilities, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MainAbility])

	if err != nil {
		return nil, err
	}

	if len(abilities) == 0 {
		return nil, fmt.Errorf("ability not found")
	}

	return &abilities[0], nil
}

func UpdateMainAbility(ctx context.Context, id int, ability models.MainAbility) error {

	query := `
	UPDATE main_abilities
	SET	name=$1,
	description=$2,
	type=$3,
	tier=$4,
	duration=$5,
	cooldown=$6,
	spike_modifier=$7,
	jump_modifier=$8,
	set_modifier=$9,
	receive_modifier=$10,
	ball_force_multiplier=$11
	WHERE id=$12
	`

	_, err := database.DB.Exec(ctx, query,
		ability.Name,
		ability.Description,
		ability.Type,
		ability.Tier,
		ability.Duration,
		ability.Cooldown,
		ability.SpikeModifier,
		ability.JumpModifier,
		ability.SetModifier,
		ability.ReceiveModifier,
		ability.BallForceMultiplier,
		id,
	)

	return err
}

func DeleteMainAbility(ctx context.Context, id int) error {

	query := `DELETE FROM main_abilities WHERE id=$1`

	_, err := database.DB.Exec(ctx, query, id)

	return err
}
