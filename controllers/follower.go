package controllers

import (
	"SocialMediaTracker/models"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type FollowerInput struct {
	Follower  uint `json:"follower" binding:"required"`
	AccountID uint `json:"AccountID" binding:"required"`
	//Handle   string `json:"handle" binding:"required"`
	//Platform string `json:"platform" binding:"required"`
}

func AddFollower(c *gin.Context) {

	var input FollowerInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	f := models.Follower{}
	f.Follower = input.Follower
	f.AccountID = input.AccountID

	_, err := f.SaveFollower()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Add follower count success"})

}

type FollowerRetrieve struct {
	AccountID uint `json:"AccountID" binding:"required"`
	//Handle   string `json:"handle" binding:"required"`
	//Platform string `json:"platform" binding:"required"`
}

func GetFollowers(c *gin.Context) {
	var input FollowerRetrieve

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	f := models.Follower{}
	f.AccountID = input.AccountID

	follower, err := models.GetFollowerByID(f.AccountID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"follower": follower})

}

func Daily(c *gin.Context) {
	// Do this all for each entry in account

	//Instagram API
	resp, err := http.Get("https://graph.facebook.com/v3.2/17841405309211844?fields=business_discovery.username(bluebottle){followers_count}")

	//Facebook API
	//resp, err := http.Get("https://graph.facebook.com/PAGE_ID/insights?metric=page_follows&access_token=ACCESS_TOKEN")

	//Twitter API
	//resp, err := http.Get("https://api.twitter.com/2/users/[ID]?user.fields=public_metrics")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": string(body)})

}
