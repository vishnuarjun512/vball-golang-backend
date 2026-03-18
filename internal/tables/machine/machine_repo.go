package machine

import (
	"context"
	"database/sql"
	"time"
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

func GetAvailablePorts_Repo(ctx context.Context, machineID int) ([]int, error) {

	// 1️⃣ Get port range
	var portStart int
	var portEnd int

	err := database.DB.QueryRow(ctx,
		`SELECT port_start, port_end FROM machines WHERE id=$1`,
		machineID,
	).Scan(&portStart, &portEnd)

	if err != nil {
		return nil, err
	}

	// 2️⃣ Get used ports
	rows, err := database.DB.Query(ctx,
		`SELECT port FROM game_servers WHERE machine_id=$1`,
		machineID,
	)
	if err != nil {
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

type Machine struct {
	ID             int       `json:"id"`
	RegionID       int       `json:"region_id"`
	IPAddress      string    `json:"ip_address"`
	CPUCores       int       `json:"cpu_cores"`
	RAMGB          int       `json:"ram_gb"`
	Status         string    `json:"status"`
	PortStart      int       `json:"port_start"`
	PortEnd        int       `json:"port_end"`
	AvailablePorts int       `json:"available_ports"`
	CreatedAt      time.Time `json:"created_at"`
}

// db is your global connection, adjust according to your project setup
var db *sql.DB

func DeleteMachine_Repo(id int) error {
	_, err := db.Exec("DELETE FROM machines WHERE id=$1", id)
	return err
}

func UpdateMachine_Repo(id int, m Machine) error {
	query := `UPDATE machines SET region_id=$1, ip_address=$2, cpu_cores=$3, ram_gb=$4, status=$5, port_start=$6, port_end=$7, available_ports=$8 WHERE id=$9`
	_, err := db.Exec(query, m.RegionID, m.IPAddress, m.CPUCores, m.RAMGB, m.Status, m.PortStart, m.PortEnd, m.AvailablePorts, id)
	return err
}

func CreateMachine_Repo(m Machine) (int, error) {
	var id int
	query := `INSERT INTO machines (region_id, ip_address, cpu_cores, ram_gb, status, port_start, port_end, available_ports) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err := db.QueryRow(query, m.RegionID, m.IPAddress, m.CPUCores, m.RAMGB, m.Status, m.PortStart, m.PortEnd, m.AvailablePorts).Scan(&id)
	return id, err
}
