package machine

import (
	"context"
	"vball/internal/models"
)

func GetAllMachines_Service() ([]models.Machine, error) {
	return GetAllMachines_Repo(context.Background())
}

func CreateMachine_Service(m Machine) (int, error) {
	// You can add validation here (e.g., check if IP is valid)
	return CreateMachine_Repo(m)
}

func UpdateMachine_Service(id int, m Machine) error {
	return UpdateMachine_Repo(id, m)
}

func DeleteMachine_Service(id int) error {
	return DeleteMachine_Repo(id)
}
