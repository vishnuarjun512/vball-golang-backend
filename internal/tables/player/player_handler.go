package player

import (
	"github.com/gin-gonic/gin"
)

func GetPlayerBySteamID_Handler(c *gin.Context) {

	steamID := c.Param("steamid")

	player, err := GetPlayerBySteamID_Service(steamID)

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
