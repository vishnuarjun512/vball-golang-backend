package vps

import (
	"net/http"
	"strconv"
	"vball/internal/services/vps"

	"github.com/gin-gonic/gin"
)

type GameServerReq struct {
	RegionID   int `json:"regionId"`
	MachineID  int `json:"machineId"`
	Port       int `json:"port"`
	MaxPlayers int `json:"maxPlayers"`
}

func CreateGameServer_Handler(c *gin.Context) {
	var req GameServerReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	gameServer, err := vps.CreateGameServer_Service(req.MachineID, req.Port, req.MaxPlayers)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to create game server",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"gameServer": gameServer,
	})
}

func GetGameServer_Handler(c *gin.Context) {
	// 1. Get the parameter as a string
	idParam := c.Param("id")

	// 2. Convert string to int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id must be an integer",
		})
		return
	}

	// 3. Call your service with the integer
	gameServer, err := vps.GetGameServer_Service(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{ // Use 404 for "Not Found"
			"error":   true,
			"message": "Failed to get game server",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":      false,
		"gameServer": gameServer,
	})
}
