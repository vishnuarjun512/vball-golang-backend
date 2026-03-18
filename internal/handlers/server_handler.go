package handlers

import (
	"vball/internal/services"
	"vball/utils"

	"github.com/gin-gonic/gin"
)

func JoinHandler(c *gin.Context) {

	type JoinRequest struct {
		PlayerID string `json:"playerId"`
		Region   string `json:"region"`
	}

	var req JoinRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, 400, "Invalid Request Body:", err)
		return
	}

	ip, port, err := services.JoinPlayer_Service(c.Request.Context(), req.PlayerID, req.Region)

	if err != nil {
		utils.SendError(c, 500, "Error joining player:", err)
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
		utils.SendError(c, 400, "Invalid Request Body:", err)
		return
	}

	err := services.LeavePlayer_Service(c.Request.Context(), req.PlayerID)

	if err != nil {
		utils.SendError(c, 500, "Failed to Leave Player Service:", err)
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
		utils.SendError(c, 400, "Invalid Request Body:", err)
		return
	}

	err := services.SyncServerPlayers_Service(
		c.Request.Context(),
		req.ServerID,
		req.Players,
	)

	if err != nil {
		utils.SendError(c, 400, "Failed to Sync Server Players:", err)
		return
	}

	c.JSON(200, gin.H{
		"status": "server players synced",
	})
}
