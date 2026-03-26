package machine

import (
	"context"
	"vball/internal/database"
)

func GetAllMachines_Repo(ctx context.Context) ([]MachineSend, error) {
	// We select into a temporary struct that matches the DB columns exactly
	query := `
		SELECT id, machine_name, region_id, ip_address, cpu_cores, ram_gb, status, port_start, port_end, created_at
		FROM machines
	`

	rows, err := database.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 1. Collect from DB. Since MachineSend has []int for ports but DB has int,
	// we use a helper to scan everything EXCEPT the available_ports first.
	var machines []MachineSend
	for rows.Next() {
		var m MachineSend
		// We do NOT scan available_ports here because the types don't match (int vs []int)
		err := rows.Scan(
			&m.ID, &m.MachineName, &m.RegionID, &m.IPAddress, &m.CPUCores,
			&m.RAMGB, &m.Status, &m.PortStart, &m.PortEnd, &m.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		machines = append(machines, m)
	}

	// 2. Compute the ACTUAL array of ports for each machine
	for i := range machines {
		ports, err := GetAvailablePorts_Repo(ctx, machines[i].ID)
		if err != nil {
			return nil, err
		}
		machines[i].AvailablePorts = ports
	}

	return machines, nil
}

func GetAvailablePorts_Repo(ctx context.Context, machineID int) ([]int, error) {
	var portStart, portEnd int

	// Get range from DB
	err := database.DB.QueryRow(ctx,
		`SELECT port_start, port_end FROM machines WHERE id=$1`,
		machineID,
	).Scan(&portStart, &portEnd)
	if err != nil {
		return nil, err
	}

	// Get ports that are already taken in game_servers table
	rows, err := database.DB.Query(ctx, `SELECT port FROM game_servers WHERE machine_id=$1`, machineID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usedPorts := make(map[int]bool)
	for rows.Next() {
		var p int
		if err := rows.Scan(&p); err != nil {
			return nil, err
		}
		usedPorts[p] = true
	}

	// Create the slice of free port numbers
	var availablePorts []int
	for port := portStart; port <= portEnd; port++ {
		if !usedPorts[port] {
			availablePorts = append(availablePorts, port)
		}
	}

	return availablePorts, nil
}

func GetMachine_Repo(ctx context.Context, id int) (MachineSend, error) {
	var m MachineSend
	query := `
		SELECT id, region_id, ip_address, cpu_cores, ram_gb, status, port_start, port_end, created_at
		FROM machines WHERE id=$1
	`
	// Again, we skip scanning the 'available_ports' column into the []int slice
	err := database.DB.QueryRow(ctx, query, id).Scan(
		&m.ID, &m.RegionID, &m.IPAddress, &m.CPUCores,
		&m.RAMGB, &m.Status, &m.PortStart, &m.PortEnd, &m.CreatedAt,
	)
	if err != nil {
		return MachineSend{}, err
	}

	// Fill the slice with real port numbers
	ports, err := GetAvailablePorts_Repo(ctx, m.ID)
	if err != nil {
		return MachineSend{}, err
	}
	m.AvailablePorts = ports

	return m, nil
}

func CreateMachine_Repo(ctx context.Context, m MachineDB) (int, error) {
	var id int
	query := `
		INSERT INTO machines (region_id, ip_address, cpu_cores, ram_gb, status, port_start, port_end, available_ports) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id
	`
	// When creating, we send the 'AvailablePorts' as an INT (the count)
	err := database.DB.QueryRow(ctx, query,
		m.RegionID, m.IPAddress, m.CPUCores, m.RAMGB, m.Status, m.PortStart, m.PortEnd, m.AvailablePorts,
	).Scan(&id)
	return id, err
}

func UpdateMachine_Repo(ctx context.Context, id int, m MachineDB) error {
	query := `
		UPDATE machines 
		SET region_id=$1, ip_address=$2, cpu_cores=$3, ram_gb=$4, status=$5, port_start=$6, port_end=$7, available_ports=$8
		WHERE id=$9
	`
	_, err := database.DB.Exec(ctx, query,
		m.RegionID, m.IPAddress, m.CPUCores, m.RAMGB, m.Status, m.PortStart, m.PortEnd, m.AvailablePorts, id,
	)
	return err
}

func DeleteMachine_Repo(ctx context.Context, id int) error {
	_, err := database.DB.Exec(ctx, "DELETE FROM machines WHERE id=$1", id)
	return err
}
