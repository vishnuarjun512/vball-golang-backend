package machine

import (
	"context"
	"fmt"
	"math/rand"
)

func GetAllMachines_Service() ([]MachineSend, error) {
	return GetAllMachines_Repo(context.Background())
}

func GetMachine_Service(id int) (MachineSend, error) {
	return GetMachine_Repo(context.Background(), id)
}

// Helper to generate a random IPv4 address
func generateFakeIP() string {
	return fmt.Sprintf("%d.%d.%d.%d",
		rand.Intn(255)+1, rand.Intn(255), rand.Intn(255), rand.Intn(255))
}

// Helper to pick a realistic RAM value
func generateRandomRAM() int {
	ramOptions := []int{4, 8, 16, 32, 64, 128}
	return ramOptions[rand.Intn(len(ramOptions))]
}

// Helper to pick a realistic CORE value
func generateRandomCORE(ram int) int {
	return rand.Intn(ram)
}

func CreateMachine_Service(m MachineCreateReq) (int, error) {
	// Translate MachineSend -> MachineDB
	// Calculate the count: (End - Start) + 1
	// Example: 7010 - 7000 = 10. +1 = 11 total ports.
	// If the request didn't provide an IP or RAM, generate them
	var ip string
	ip = generateFakeIP()

	var ram int
	ram = generateRandomRAM()

	var core int
	core = generateRandomCORE(ram)

	dbModel := MachineDB{
		MachineName:    m.MachineName,
		RegionID:       m.RegionID,
		IPAddress:      ip,
		CPUCores:       core,
		RAMGB:          ram,
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
