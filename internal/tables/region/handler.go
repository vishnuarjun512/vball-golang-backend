package region

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegionReq struct {
	RegionName string `json:"regionName"`
}

func GetAllRegions_Hanlder(c *gin.Context) {
	regions, err := GetRegions_Service()
	if err != nil {
		fmt.Println("Error getting all Regions:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get Regions",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"regions": regions,
	})
}

func CreateRegion_Handler(c *gin.Context) {
	var req RegionReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	id, err := CreateRegion_Service(req.RegionName)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create Region",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"id":    id,
	})
}

func UpdateRegion_Handler(c *gin.Context) {

	regionId := c.Param("regionId")

	var req RegionReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	err := UpdateRegion_Service(req.RegionName, regionId)

	if err != nil {
		fmt.Println("Error updating Region:", err)
		c.JSON(500, gin.H{"message": "update failed", "error": true})
		return
	}

	c.JSON(200, gin.H{"message": "updated", "error": false})
}

func DeleteRegion_Handler(c *gin.Context) {
	id := c.Param("id")

	err := DeleteRegion_Service(id)

	if err != nil {
		fmt.Println("Error deleting Region:", err)
		c.JSON(500, gin.H{"error": true, "message": "deletion failed"})
		return
	}

	fmt.Println("Region deleted successfully")
	c.JSON(200, gin.H{"message": "deleted", "error": false})
}
