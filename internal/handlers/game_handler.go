package handlers

import (
	"net/http"

	"vball/utils"

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
		utils.SendError(c, 500, "Error fetching players:", err)
		return
	}

	mainAbilities, subAbilities, err := services.GetAllAbilities_Service()
	if err != nil {
		utils.SendError(c, 500, "Error fetching Abilites", err)
		return
	}

	regions, err := region.GetRegions_Service()
	if err != nil {
		utils.SendError(c, 500, "Error fetching Regions", err)
		return
	}

	machines, err := machine.GetAllMachines_Service()
	if err != nil {
		utils.SendError(c, 500, "Error fetching Machines:", err)
		return
	}

	gameServers, err := gameserver.GetAllGameServers_Service()
	if err != nil {
		utils.SendError(c, 500, "Error fetching Gameservers:", err)
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
		utils.SendError(c, 500, "Failed to load abilities:", err)
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
	var req SteamLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, 500, "Invalid request body", err)
		return
	}

	redirectURL, err := player.GetSteamLogin_Service(req.SteamID, req.Username)
	if err != nil {
		utils.SendError(c, 500, "Failed to get steam login URL:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"redirectURL": redirectURL,
	})
}
