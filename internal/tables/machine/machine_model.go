package machine

import "time"

type MachineDB struct {
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

type MachineSend struct {
	ID             int       `json:"id"`
	RegionID       int       `json:"region_id"`
	IPAddress      string    `json:"ip_address"`
	CPUCores       int       `json:"cpu_cores"`
	RAMGB          int       `json:"ram_gb"`
	Status         string    `json:"status"`
	PortStart      int       `json:"port_start"`
	PortEnd        int       `json:"port_end"`
	AvailablePorts []int     `json:"available_ports"`
	CreatedAt      time.Time `json:"created_at"`
}
