package gameserver

import (
	"context"
	"fmt"
	"vball/internal/database"
	"vball/internal/models"

	"github.com/jackc/pgx/v5"
)

// Changed machine_id to int to match database integer column types
func CreateGameServer_Repo(ctx context.Context, machine_id int, port int, max_players int) (*models.GameServer, error) {

	query := `
	INSERT INTO game_servers (machine_id, port, max_players)
	VALUES ($1, $2, $3)
	RETURNING id`

	var id int
	err := database.DB.QueryRow(ctx, query, machine_id, port, max_players).Scan(&id)
	if err != nil {
		return nil, err
	}

	return GetGameServer_Repo(ctx, id)
}

// Function signature updated to return (*models.GameServer, error)
func GetGameServer_Repo(ctx context.Context, id int) (*models.GameServer, error) {
	query := `
	SELECT id, machine_id, port, max_players, current_players, status, uptime, created_at
	FROM game_servers
	WHERE id = $1`

	rows, err := database.DB.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Use pgx.CollectOneRow for single record lookups
	server, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.GameServer])
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("game server not found")
		}
		return nil, err
	}

	return &server, nil
}

func DeleteGameServer_Repo(ctx context.Context, id int) error {
	query := `DELETE FROM game_servers WHERE id = $1`

	result, err := database.DB.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	// Check if any row was actually deleted
	if result.RowsAffected() == 0 {
		return fmt.Errorf("game server with id %d not found", id)
	}

	return nil
}

func UpdateGameServer_Repo(ctx context.Context, id int, status string, currentPlayers int) (*models.GameServer, error) {
	query := `
	UPDATE game_servers 
	SET status = $1, current_players = $2
	WHERE id = $3
	RETURNING id, machine_id, port, max_players, current_players, status, uptime, created_at`

	var server models.GameServer

	// Since we are returning the record, we use QueryRow
	err := database.DB.QueryRow(ctx, query, status, currentPlayers, id).Scan(
		&server.ID,
		&server.MachineID,
		&server.Port,
		&server.MaxPlayers,
		&server.CurrentPlayers,
		&server.Status,
		&server.Uptime,
		&server.CreatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("game server not found")
		}
		return nil, err
	}

	return &server, nil
}

func GetAllGameServers_Repo(ctx context.Context) ([]models.GameServer, error) {

	query := `
		SELECT 
			id,
			machine_id,
			port,
			max_players,
			current_players,
			status,
			uptime,
			created_at
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
