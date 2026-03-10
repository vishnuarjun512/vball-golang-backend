package repositories

import (
	"context"
	"vball/internal/database"
	"vball/internal/models"
)

func CreateSubAbility(ctx context.Context, ability models.SubAbility) error {

	query := `
	INSERT INTO sub_abilities
	(name, description, tier, modifier_type, modifier_value)
	VALUES ($1,$2,$3,$4,$5)
	`

	_, err := database.DB.Exec(ctx, query,
		ability.Name,
		ability.Description,
		ability.Tier,
		ability.ModifierType,
		ability.ModifierValue,
	)

	return err
}

func GetSubAbilities(ctx context.Context) ([]models.SubAbility, error) {

	query := `SELECT id,name,description,tier,modifier_type,modifier_value FROM sub_abilities`

	rows, err := database.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var abilities []models.SubAbility

	for rows.Next() {

		var ability models.SubAbility

		err := rows.Scan(
			&ability.ID,
			&ability.Name,
			&ability.Description,
			&ability.Tier,
			&ability.ModifierType,
			&ability.ModifierValue,
		)

		if err != nil {
			return nil, err
		}

		abilities = append(abilities, ability)
	}

	return abilities, nil
}

func GetSubAbility(ctx context.Context, id int) (*models.SubAbility, error) {

	query := `
	SELECT id,name,description,tier,modifier_type,modifier_value
	FROM sub_abilities
	WHERE id=$1
	`

	row := database.DB.QueryRow(ctx, query, id)

	var ability models.SubAbility

	err := row.Scan(
		&ability.ID,
		&ability.Name,
		&ability.Description,
		&ability.Tier,
		&ability.ModifierType,
		&ability.ModifierValue,
	)

	if err != nil {
		return nil, err
	}

	return &ability, nil
}

func UpdateSubAbility(ctx context.Context, id int, ability models.SubAbility) error {

	query := `
	UPDATE sub_abilities
	SET modifier_type=$1,
	    modifier_value=$2
	WHERE id=$3
	`

	_, err := database.DB.Exec(ctx, query,
		ability.ModifierType,
		ability.ModifierValue,
		id,
	)

	return err
}

func DeleteSubAbility(ctx context.Context, id int) error {

	query := `DELETE FROM sub_abilities WHERE id=$1`

	_, err := database.DB.Exec(ctx, query, id)

	return err
}
