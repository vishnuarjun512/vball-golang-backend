package player

import (
	"net/http"
	"vball/utils"

	"github.com/gin-gonic/gin"
)

func GetPlayerBySteamID_Handler(c *gin.Context) {
	steamID := c.Param("steamid")
	if steamID == "" {
		utils.SendError(c, http.StatusBadRequest, "SteamID is required", nil)
		return
	}

	player, err := GetPlayerBySteamID_Service(steamID)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to fetch player", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "player": player})
}

// Added this for your /game/auth route
func GetSteamLogin_Handler(c *gin.Context) {
	var req struct {
		SteamID  string `json:"steamId"`
		Username string `json:"username"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid login data", err)
		return
	}

	player, err := GetSteamLogin_Service(req.SteamID, req.Username)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Login/Registration failed", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "player": player, "message": "Login successful"})
}

func CreatePlayer_Handler(c *gin.Context) {
	var req struct {
		SteamID  string `json:"steamId"`
		Username string `json:"username"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid login data", err)
		return
	}

	playerId, err := CreatePlayer_Service(req.SteamID, req.Username)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Player Creation Failed", err)
		return
	}

	player, err := GetPlayerBySteamID_Service(playerId)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to Fetch Player after Creation", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "player": player, "message": "Player Created Successfully"})
}

func DeletePlayer_Handler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.SendError(c, http.StatusBadRequest, "Invalid login data", nil)
		return
	}

	err := DeletePlayer_Service(id)

	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Error Deleting Player", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Player Deleted Successfully",
	})

}
