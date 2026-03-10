package handlers

import (
	"github.com/gin-gonic/gin"

	"vball/internal/services"
)

func GetAllPlayersLoadOut_Handler(c *gin.Context) {

	players, err := services.GetAllPlayersLoadOut_Service()

	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to load players",
		})
		return
	}

	c.JSON(200, players)
}

func GetPlayerBySteamID_Handler(c *gin.Context) {

	steamID := c.Param("steamid")

	player, err := services.GetPlayerBySteamID_Service(steamID)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to fetch player",
		})
		return
	}

	if player == nil {
		c.JSON(404, gin.H{
			"error": "player not found",
		})
		return
	}

	c.JSON(200, player)
}
