package region

import (
	"context"
	"fmt"
	"vball/internal/database"
	"vball/internal/models"

	"github.com/jackc/pgx/v5"
)

func GetAllMachines_Repo(ctx context.Context) ([]models.Machine, error) {

	query := `
		SELECT 
			id,
			region_id,
			ip_address,
			cpu_cores,
			ram_gb,
			status,
			port_start,
			port_end
		FROM machines
	`

	rows, err := database.DB.Query(ctx, query)
	if err != nil {
		fmt.Println("VPS_REPO_ERROR: Cannot find any Machines")
		return nil, err
	}
	defer rows.Close()

	machines, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Machine])
	if err != nil {
		return nil, err
	}

	// compute available ports for each machine
	for i := range machines {

		availablePorts, err := GetAvailablePorts_Repo(ctx, machines[i].ID)
		if err != nil {
			return nil, err
		}

		machines[i].AvailablePorts = availablePorts
	}

	return machines, nil
}
