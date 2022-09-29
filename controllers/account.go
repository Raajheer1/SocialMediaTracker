package controllers

import (
	"SocialMediaTracker/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NewAccountInput struct {
	Handle     string `json:"handle" binding:"required"`
	Platform   string `json:"platform" binding:"required"`
	Department string `json:"department" binding:"required"`
}

func RegisterAccount(c *gin.Context) {

	var input NewAccountInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a := models.Account{}

	a.Handle = input.Handle
	a.Department = input.Department
	a.Platform = input.Platform

	_, err := a.SaveAccount()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

}
