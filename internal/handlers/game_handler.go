package handlers

import (
	"fmt"
	"net/http"
	"vball/internal/services"
	"vball/internal/tables/gameserver"
	"vball/internal/tables/machine"
	"vball/internal/tables/player"
	"vball/internal/tables/region"

	"github.com/gin-gonic/gin"
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

	mainAbilities, subAbilities, err := services.GetAllAbilities_Service()

	if err != nil {
		fmt.Println("Error fetching abilities:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to load abilities",
		})
		return
	}

	regions, err := region.GetRegions_Service()
	if err != nil {
		fmt.Println("Error fetching regions:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to load regions",
		})
		return
	}

	machines, err := machine.GetAllMachines_Service()
	if err != nil {
		fmt.Println("Error fetching Machines:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to load Machines",
		})
		return
	}

	gameServers, err := gameserver.GetAllGameServers_Service()
	if err != nil {
		fmt.Println("Error fetching Game Servers:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to load Game Servers",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"players":       players,
		"mainAbilities": mainAbilities,
		"subAbilities":  subAbilities,
		"regions":       regions,
		"machines":      machines,
		"gameServers":   gameServers,
	})

}

func GetGameAbilities(c *gin.Context) {

	mainAbilities, subAbilities, err := services.GetAllAbilities_Service()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to load abilities",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mainAbilities": mainAbilities,
		"subAbilities":  subAbilities,
	})
}

type SteamLoginRequest struct {
	SteamID  string `json:"steamId"`
	Username string `json:"username"`
}

func GetSteamLogin_Handler(c *gin.Context) {
	// bind incoming JSON into the request struct
	var req SteamLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	redirectURL, err := player.GetSteamLogin_Service(req.SteamID, req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get steam login URL",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"redirectURL": redirectURL,
	})
}
