package region

import (
	"fmt"
	"net/http"
	region "vball/internal/services/regions"

	"github.com/gin-gonic/gin"
)

type MachineReq struct {
	RegionId string `json:"region_id"`
}

func GetAllMachines_Handler(c *gin.Context) {
	machines, err := region.GetAllMachines_Service()
	if err != nil {
		fmt.Println("Error getting all machines:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get machines",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":    false,
		"machines": machines,
	})
}
