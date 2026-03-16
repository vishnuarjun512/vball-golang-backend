package region

import (
	"context"
	"fmt"
	"vball/internal/database"
	"vball/internal/models"

	"github.com/jackc/pgx/v5"
)

func GetAllGameServers_Repo(ctx context.Context) ([]models.GameServer, error) {

	query := `
		SELECT 
			id,
			machine_id,
			port,
			max_players,
			current_players,
			status,
			uptime
		FROM game_servers
	`

	rows, err := database.DB.Query(ctx, query)
	if err != nil {
		fmt.Println("VPS_REPO_ERROR: Cannot find any Game Servers", err)
		return nil, err
	}
	defer rows.Close()

	gameServers, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.GameServer])
	if err != nil {
		return nil, err
	}

	return gameServers, nil
}
