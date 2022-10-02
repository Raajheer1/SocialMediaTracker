package controllers

import (
	"SocialMediaTracker/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

type FollowerInput struct {
	Follower  uint `json:"follower" binding:"required"`
	AccountID uint `json:"AccountID" binding:"required"`
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

type IGResponse struct {
	BusinessDiscovery struct {
		FollowerCount uint `json:"followers_count"`
	} `json:"business_discovery"`
}

func Daily(c *gin.Context) {
	var update []FollowerInput

	accounts, err := models.GetAllAccounts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	AccessToken := os.Getenv("IG_ACCESS_TOKEN")

	for _, account := range accounts {
		if account.Platform == "IG" {
			resp, err := http.Get(fmt.Sprintf("https://graph.facebook.com/v15.0/17841454380310710?fields=business_discovery.username(%s){followers_count,media_count}&access_token=%s", account.Handle, AccessToken))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}

			var res IGResponse
			if err := json.Unmarshal(body, &res); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}

			//c.JSON(http.StatusOK, gin.H{"data": string(body)})
			followerEntry := FollowerInput{
				Follower:  res.BusinessDiscovery.FollowerCount,
				AccountID: account.ID,
			}
			update = append(update, followerEntry)
		}
	}

	for _, newFollower := range update {

		f := models.Follower{}
		f.Follower = newFollower.Follower
		f.AccountID = newFollower.AccountID

		_, err := f.SaveFollower()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}

	//fmt.Println(update)

	//Facebook API
	//resp, err := http.Get("https://graph.facebook.com/PAGE_ID/insights?metric=page_follows&access_token=ACCESS_TOKEN")

	//Twitter API
	//resp, err := http.Get("https://api.twitter.com/2/users/[ID]?user.fields=public_metrics")

	c.JSON(http.StatusOK, gin.H{"data": "Follower count updated successfully."})

}
