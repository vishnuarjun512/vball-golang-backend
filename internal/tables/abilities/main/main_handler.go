package mainAbility

import (
	"net/http"
	"strconv"

	"vball/utils" // Import your new helper

	"github.com/gin-gonic/gin"
)

func CreateMainAbility(c *gin.Context) {
	var newAbility CreateAbilityRequest

	if err := c.ShouldBindJSON(&newAbility); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	createdAbility, err := CreateMainAbility_Service(newAbility)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Error creating main ability", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"error":   false,
		"message": "Main ability created successfully",
		"ability": createdAbility,
	})
}

func GetMainAbilities(c *gin.Context) {
	abilities, err := GetMainAbilities_Service()
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Error fetching main abilities", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":     false,
		"message":   "Main abilities fetched successfully",
		"abilities": abilities,
	})
}

func GetMainAbility(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		utils.SendError(c, http.StatusBadRequest, "Invalid Ability ID", nil)
		return
	}

	ability, err := GetMainAbility_Service(id)
	if err != nil {
		utils.SendError(c, http.StatusNotFound, "Main ability not found", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"ability": ability,
	})
}

func UpdateMainAbility(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		utils.SendError(c, http.StatusBadRequest, "Invalid Ability ID", nil)
		return
	}

	var ability MainAbility
	if err := c.ShouldBindJSON(&ability); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid JSON input", err)
		return
	}

	err = UpdateMainAbility_Service(id, ability)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Error updating main ability", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Main ability updated successfully",
	})
}

func DeleteMainAbility(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		utils.SendError(c, http.StatusBadRequest, "Invalid Ability ID", nil)
		return
	}

	err = DeleteMainAbility_Service(id)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Error deleting main ability", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Main ability deleted successfully",
	})
}
