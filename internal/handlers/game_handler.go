package handlers

import (
	"net/http"
	"vball/internal/services"

	"github.com/gin-gonic/gin"
)

func GetGameAbilities(c *gin.Context) {

	mainAbilities, subAbilities, err := services.GetAllAbilities()

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

	redirectURL, err := services.GetSteamLogin_Service(req.SteamID, req.Username)
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
