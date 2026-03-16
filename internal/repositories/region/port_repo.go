package region

import (
	"context"
	"fmt"
	"vball/internal/database"
)

func GetAvailablePorts_Repo(ctx context.Context, machineID int) ([]int, error) {

	// 1️⃣ Get port range
	var portStart int
	var portEnd int

	err := database.DB.QueryRow(ctx,
		`SELECT port_start, port_end FROM machines WHERE id=$1`,
		machineID,
	).Scan(&portStart, &portEnd)

	if err != nil {
		fmt.Println("PORT_REPO_ERROR: Cannot fetch machine port range")
		return nil, err
	}

	// 2️⃣ Get used ports
	rows, err := database.DB.Query(ctx,
		`SELECT port FROM game_servers WHERE machine_id=$1`,
		machineID,
	)
	if err != nil {
		fmt.Println("PORT_REPO_ERROR: Cannot fetch used ports")
		return nil, err
	}
	defer rows.Close()

	usedPorts := map[int]bool{}

	for rows.Next() {
		var port int
		if err := rows.Scan(&port); err != nil {
			return nil, err
		}
		usedPorts[port] = true
	}

	// 3️⃣ Generate available ports
	var availablePorts []int

	for port := portStart; port < portEnd; port++ {
		if !usedPorts[port] {
			availablePorts = append(availablePorts, port)
		}
	}

	return availablePorts, nil
}
