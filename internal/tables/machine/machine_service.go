package machine

import (
	"context"
)

func GetAllMachines_Service() ([]MachineSend, error) {
	return GetAllMachines_Repo(context.Background())
}

func GetMachine_Service(id int) (MachineSend, error) {
	return GetMachine_Repo(context.Background(), id)
}

func CreateMachine_Service(m MachineSend) (int, error) {
	// Translate MachineSend -> MachineDB
	// Calculate the count: (End - Start) + 1
	// Example: 7010 - 7000 = 10. +1 = 11 total ports.
	dbModel := MachineDB{
		RegionID:       m.RegionID,
		IPAddress:      m.IPAddress,
		CPUCores:       m.CPUCores,
		RAMGB:          m.RAMGB,
		Status:         m.Status,
		PortStart:      m.PortStart,
		PortEnd:        m.PortEnd,
		AvailablePorts: (m.PortEnd - m.PortStart) + 1,
	}

	return CreateMachine_Repo(context.Background(), dbModel)
}

func UpdateMachine_Service(id int, m MachineSend) error {
	// Translate MachineSend -> MachineDB
	dbModel := MachineDB{
		RegionID:       m.RegionID,
		IPAddress:      m.IPAddress,
		CPUCores:       m.CPUCores,
		RAMGB:          m.RAMGB,
		Status:         m.Status,
		PortStart:      m.PortStart,
		PortEnd:        m.PortEnd,
		AvailablePorts: (m.PortEnd - m.PortStart) + 1,
	}

	return UpdateMachine_Repo(context.Background(), id, dbModel)
}

func DeleteMachine_Service(id int) error {
	return DeleteMachine_Repo(context.Background(), id)
}
