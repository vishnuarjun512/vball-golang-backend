package region

import (
	"context"
	"vball/internal/database"
	"vball/internal/models"

	"github.com/jackc/pgx/v5"
)

func GetAllRegions_Repo(ctx context.Context) ([]models.Region, error) {

	query := `SELECT id,region_name,region_code from regions`

	rows, err := database.DB.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	regions, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Region])

	return regions, err
}

func CreateRegion_Repo(ctx context.Context, regionName string, regionCode string) (string, error) {

	query := `
	INSERT INTO regions (region_name, region_code)
	VALUES ($1,$2)
	RETURNING id
	`

	var regionId string
	err := database.DB.QueryRow(ctx, query, regionName, regionCode).Scan(&regionId)

	if err != nil {
		return "", err
	}

	return regionId, nil
}

func UpdateRegion_Repo(ctx context.Context, regionId string, regionName string) error {

	query := `
		UPDATE regions
		SET name=$1
		WHERE id=$2
	`

	_, err := database.DB.Exec(ctx, query, regionName, regionId)

	return err
}

func DeleteRegion_Repo(ctx context.Context, id string) error {

	query := `DELETE FROM regions WHERE id=$1`

	_, err := database.DB.Exec(ctx, query, id)

	return err
}
