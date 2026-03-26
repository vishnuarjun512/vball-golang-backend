package machine

import (
	"strconv"
	"vball/utils"

	"github.com/gin-gonic/gin"
)

func GetAllMachines_Handler(c *gin.Context) {
	machines, err := GetAllMachines_Service()
	if err != nil {
		utils.SendError(c, 500, "Failed to fetch machines", err)
		return
	}
	c.JSON(200, gin.H{"error": false, "machines": machines})
}

type MachineCreateReq struct {
	ID          int    `json:"id"`
	MachineName string `json:"machine_name"`
	RegionID    int    `json:"region_id"`
	Status      string `json:"status"`
	PortStart   int    `json:"port_start"`
	PortEnd     int    `json:"port_end"`
}

func CreateMachine_Handler(c *gin.Context) {
	var m MachineCreateReq // Frontend sends the port range
	if err := c.ShouldBindJSON(&m); err != nil {
		utils.SendError(c, 400, "Invalid input", err)
		return
	}

	id, err := CreateMachine_Service(m)
	if err != nil {
		utils.SendError(c, 500, "Creation failed", err)
		return
	}
	machine, err := GetMachine_Service(id)
	c.JSON(201, gin.H{"error": false, "id": id, "message": "Machine Created", "machine": machine})
}

func UpdateMachine_Handler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var m MachineSend
	if err := c.ShouldBindJSON(&m); err != nil {
		utils.SendError(c, 400, "Invalid input", err)
		return
	}

	if err := UpdateMachine_Service(id, m); err != nil {
		utils.SendError(c, 500, "Update failed", err)
		return
	}
	c.JSON(200, gin.H{"error": false, "message": "Machine Updated"})
}

func DeleteMachine_Hander(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteMachine_Service(id); err != nil {
		utils.SendError(c, 500, "Delete failed", err)
		return
	}
	c.JSON(200, gin.H{"error": false, "message": "Machine Deleted"})
}

func GetMachine_Handler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	machine, err := GetMachine_Service(id)
	if err != nil {
		utils.SendError(c, 500, "Machine not found", err)
		return
	}
	c.JSON(200, gin.H{"error": false, "message": "Machine Found", "machine": machine})
}
