package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"vball/internal/services"
	vps "vball/internal/services/vps"
)

func GetAdminLoadOut_Handler(c *gin.Context) {

	players, err := services.GetAdminLoadOut_Service()

	if err != nil {
		fmt.Println("Error fetching players:", err)
		c.JSON(500, gin.H{
			"error": "failed to load players",
		})
		return
	}

	mainAbilities, subAbilities, err := services.GetAllAbilities()

	if err != nil {
		fmt.Println("Error fetching abilities:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to load abilities",
		})
		return
	}

	regions, err := vps.GetRegions_Service()
	if err != nil {
		fmt.Println("Error fetching regions:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to load regions",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"players":       players,
		"mainAbilities": mainAbilities,
		"subAbilities":  subAbilities,
		"regions":       regions,
	})

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
