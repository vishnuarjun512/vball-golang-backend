package region

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"vball/utils"
)

func GetAllRegions_Hanlder(c *gin.Context) {
	regions, err := GetRegions_Service()
	if err != nil {

		utils.SendError(c, http.StatusBadRequest, "Error getting all Regions", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"regions": regions,
		"message": "Fetched All Regions",
	})
}

func CreateRegion_Handler(c *gin.Context) {
	var req RegionReq

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid Request Body", err)
		return
	}

	id, err := CreateRegion_Service(req.RegionName, req.RegionCode)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to Create Region", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"error":   false,
		"id":      id,
		"message": "Created Region Successfully",
	})
}

func UpdateRegion_Handler(c *gin.Context) {
	regionId := c.Param("regionId")
	if regionId == "" {
		utils.SendError(c, http.StatusBadRequest, "RegionID received Empty", nil)
		return
	}

	var req RegionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid Request Body", err)
		return
	}

	err := UpdateRegion_Service(req.RegionName, regionId)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Error updating Region", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Updated Region Successfully",
	})
}

func DeleteRegion_Handler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.SendError(c, http.StatusBadRequest, "RegionID received Empty", nil)
		return
	}

	err := DeleteRegion_Service(id)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Error deleting Region", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Deleted Region Successfully",
	})
}

func GetRegion_Handler(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		utils.SendError(c, http.StatusBadRequest, "RegionID received Empty", nil)
		return
	}

	region, err := GetRegion_Service(id)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Error Finding Region", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Found Region Successfully",
		"region":  region,
	})

}
