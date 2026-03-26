package models

import "time"

type PlayerLoadout struct {
	PlayerID string `json:"playerId"`
	SteamID  string `json:"steamId"`
	Username string `json:"username"`
	Kash     int    `json:"kash"`

	MainAbilityID *int `json:"mainAbilityId"`

	SubAbilitySlot1 *int `json:"subAbilitySlot1"`
	SubAbilitySlot2 *int `json:"subAbilitySlot2"`
	SubAbilitySlot3 *int `json:"subAbilitySlot3"`
}

type Machine struct {
	ID             int    `db:"id" json:"id"`
	RegionID       string `db:"region_id" json:"region_id"`
	IPAddress      string `db:"ip_address" json:"ip_address"`
	CPUCores       int    `db:"cpu_cores" json:"cpu_cores"`
	RamGB          int    `db:"ram_gb" json:"ram_gb"`
	Status         string `db:"status" json:"status"`
	PortStart      int    `db:"port_start" json:"port_start"`
	PortEnd        int    `db:"port_end" json:"port_end"`
	AvailablePorts []int  `db:"-" json:"available_ports"`
}

type GameServer struct {
	ID             int       `db:"id" json:"id"`
	MachineID      int       `db:"machine_id" json:"machine_id"`
	Port           int       `db:"port" json:"port"`
	MaxPlayers     int       `db:"max_players" json:"max_players"`
	CurrentPlayers int       `db:"current_players" json:"current_players"`
	Status         string    `db:"status" json:"status"`
	Uptime         time.Time `db:"uptime" json:"uptime"`
	CreatedAt      time.Time `db:"created_at" json:"createdAt"`
}

type Match struct {
	ID       string `db:"id"`
	ServerID string `db:"server_id"`
	Status   string `db:"status"`
}

type MatchPlayer struct {
	ID       string `db:"id"`
	MatchID  string `db:"match_id"`
	PlayerID string `db:"player_id"`
	Team     int    `db:"team"`
}
