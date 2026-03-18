package machine

import (
	"net/http"
	"strconv"
	"vball/utils"

	"github.com/gin-gonic/gin"
)

func GetAllMachines_Handler(c *gin.Context) {

	machines, err := GetAllMachines_Service()

	if err != nil {

		utils.SendError(c, http.StatusBadRequest, "Error getting all machines", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":    false,
		"message":  "Fetched Machines Successfully",
		"machines": machines,
	})
}

func Create(c *gin.Context) {
	var m Machine

	if err := c.ShouldBindJSON(&m); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid Body Request", err)
		return
	}

	id, err := CreateMachine_Service(m)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Failed to Create Machine", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"error": false, "id": id, "message": "Machine Created"})
}

func Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		utils.SendError(c, http.StatusBadRequest, "Invalid Machine ID", nil)
		return
	}

	var m Machine
	if err := c.ShouldBindJSON(&m); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid JSON input", err)
		return
	}

	if err := UpdateMachine_Service(id, m); err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to update machine in database", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Machine Updated Successfully"})
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := DeleteMachine_Service(id); err != nil {
		// Call it using the package name prefix: utils.SendError
		utils.SendError(c, http.StatusInternalServerError, "Delete Machine failed", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Machine Deleted"})
}
