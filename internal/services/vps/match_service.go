package vps

import (
	"context"
	"fmt"
	"vball/internal/database"

	"vball/internal/repositories/vps"
)

func JoinPlayer_Service(ctx context.Context, playerID string, region string) (string, int, error) {

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return "", 0, err
	}

	server, ip, err := vps.FindAvailableServerTx(ctx, tx, region)
	if err != nil {
		tx.Rollback(ctx)
		return "", 0, err
	}

	if server == nil {
		tx.Rollback(ctx)
		return "", 0, fmt.Errorf("no available servers")
	}

	err = vps.AddPlayerTx(ctx, tx, server.ID, playerID)
	if err != nil {
		tx.Rollback(ctx)
		return "", 0, err
	}

	err = vps.IncrementPlayersTx(ctx, tx, server.ID)
	if err != nil {
		tx.Rollback(ctx)
		return "", 0, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return "", 0, err
	}

	return ip, server.Port, nil
}

func LeavePlayer_Service(ctx context.Context, playerID string) error {

	serverID, err := vps.FindPlayerServer(ctx, playerID)
	if err != nil {
		return err
	}

	err = vps.RemovePlayer(ctx, playerID)
	if err != nil {
		return err
	}

	err = vps.DecrementPlayers(ctx, serverID)
	if err != nil {
		return err
	}

	return nil
}

func SyncServerPlayers_Service(ctx context.Context, serverID int, players []string) error {

	if serverID == 0 {
		return fmt.Errorf("invalid server id")
	}

	err := vps.SyncServerPlayers_Repo(ctx, serverID, players)
	if err != nil {
		return err
	}

	return nil
}
