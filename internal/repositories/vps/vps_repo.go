package vps

import (
	"context"
	"fmt"
	"vball/internal/database"
	"vball/internal/models"

	"github.com/jackc/pgx/v5"
)

func GetAllRegions_Repo(ctx context.Context) ([]models.Region, error) {

	query := `SELECT id,name from regions`

	rows, err := database.DB.Query(ctx, query)

	if err != nil {
		fmt.Println("VPS_REPO_ERROR: Cannot find any Regions")
		return nil, err
	}

	defer rows.Close()

	regions, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Region])

	return regions, err
}
