package vps

import (
	vps "vball/internal/services/vps"

	"fmt"

	"github.com/gin-gonic/gin"
)

func JoinHandler(c *gin.Context) {

	type JoinRequest struct {
		PlayerID string `json:"playerId"`
		Region   string `json:"region"`
	}

	var req JoinRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, err)
		return
	}

	ip, port, err := vps.JoinPlayer_Service(c.Request.Context(), req.PlayerID, req.Region)

	if err != nil {
		fmt.Printf("Error joining player: %v\n", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"serverIp": ip,
		"port":     port,
	})
}

type LeaveRequest struct {
	PlayerID string `json:"playerId"`
}

func LeaveHandler(c *gin.Context) {

	var req LeaveRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := vps.LeavePlayer_Service(c.Request.Context(), req.PlayerID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "player removed",
	})
}

type SyncRequest struct {
	ServerID int      `json:"serverId"`
	Players  []string `json:"players"`
}

func SyncServer_Handler(c *gin.Context) {

	var req SyncRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := vps.SyncServerPlayers_Service(
		c.Request.Context(),
		req.ServerID,
		req.Players,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "server players synced",
	})
}
