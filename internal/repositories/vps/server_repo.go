package vps

import (
	"context"

	"vball/internal/database"
	"vball/internal/models"

	"github.com/jackc/pgx/v5"
)

func FindAvailableServerTx(ctx context.Context, tx pgx.Tx, region string) (*models.GameServer, string, error) {

	query := `
	SELECT 
		gs.id,
		gs.machine_id,
		gs.port,
		gs.max_players,
		gs.current_players,
		m.ip_address
	FROM game_servers gs
	JOIN machines m ON m.id = gs.machine_id
	JOIN regions r ON r.id = m.region_id
	WHERE LOWER(r.name)=LOWER($1)
	AND gs.current_players < gs.max_players
	AND gs.status='running'
	ORDER BY gs.current_players ASC
	FOR UPDATE SKIP LOCKED
	LIMIT 1
	`

	var server models.GameServer
	var ip string

	err := tx.QueryRow(ctx, query, region).Scan(
		&server.ID,
		&server.MachineID,
		&server.Port,
		&server.MaxPlayers,
		&server.CurrentPlayers,
		&ip,
	)

	if err != nil {
		return nil, "", err
	}

	return &server, ip, nil
}

func SyncServerPlayers_Repo(ctx context.Context, serverID int, players []string) error {

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, `
	DELETE FROM server_players
	WHERE server_id=$1
	`, serverID)

	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	for _, player := range players {

		_, err = tx.Exec(ctx, `
		INSERT INTO server_players(server_id, player_id)
		VALUES ($1,$2)
		`, serverID, player)

		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}

	_, err = tx.Exec(ctx, `
	UPDATE game_servers
	SET current_players=$1
	WHERE id=$2
	`, len(players), serverID)

	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}

func AddPlayerTx(ctx context.Context, tx pgx.Tx, serverID int, playerID string) error {

	query := `
	INSERT INTO server_players (server_id, player_id)
	VALUES ($1,$2)
	ON CONFLICT DO NOTHING
	`

	_, err := tx.Exec(ctx, query, serverID, playerID)

	return err
}

func IncrementPlayersTx(ctx context.Context, tx pgx.Tx, serverID int) error {

	query := `
	UPDATE game_servers
	SET current_players = current_players + 1
	WHERE id=$1
	`

	_, err := tx.Exec(ctx, query, serverID)

	return err
}

func DecrementPlayers(ctx context.Context, serverID int) error {

	query := `
	UPDATE game_servers
	SET current_players = current_players - 1
	WHERE id=$1 AND current_players > 0
	`

	_, err := database.DB.Exec(ctx, query, serverID)

	return err
}

func FindPlayerServer(ctx context.Context, playerID string) (int, error) {

	query := `
	SELECT server_id
	FROM server_players
	WHERE player_id=$1
	`

	var serverID int
	err := database.DB.QueryRow(ctx, query, playerID).Scan(&serverID)

	return serverID, err
}

func RemovePlayer(ctx context.Context, playerID string) error {

	query := `
	DELETE FROM server_players
	WHERE player_id=$1
	`

	_, err := database.DB.Exec(ctx, query, playerID)

	return err
}
