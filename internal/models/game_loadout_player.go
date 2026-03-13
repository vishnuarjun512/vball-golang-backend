package models

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

type Region struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type Machine struct {
	ID        string `db:"id"`
	RegionID  string `db:"region_id"`
	Name      string `db:"name"`
	IPAddress string `db:"ip_address"`
	SSHPort   int    `db:"ssh_port"`
	Status    string `db:"status"`
}

type GameServer struct {
	ID             int    `db:"id"`
	MachineID      string `db:"machine_id"`
	Port           int    `db:"port"`
	MaxPlayers     int    `db:"max_players"`
	CurrentPlayers int    `db:"current_players"`
	Status         string `db:"status"`
	ProcessID      int    `db:"process_id"`
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
