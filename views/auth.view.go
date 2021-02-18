package views

import (
	"context"
	"fmt"
	"log"
	"negigo/models"
	"net/http"
	"strconv"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

type AuthViews struct{}

func (a AuthViews) ListAllUsers(c *gin.Context) {
	ctx := context.TODO()
	client := models.GetFirebaseAuthClient(c)
	// Note, behind the scenes, the Users() iterator will retrive 1000 Users at a time through the API
	iter := client.Users(ctx, "")
	var usersList []models.Member
	for {
		user, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("error listing users: %s\n", err)
		}
		member := models.Member{}
		member.UID = user.UID
		member.DisplayName = user.DisplayName
		if user.CustomClaims["role"] != nil {
			_role, err := strconv.Atoi(fmt.Sprintf("%v", user.CustomClaims["role"]))
			if err != nil {
				log.Println("role to int err: ", err)
				continue
			}
			member.Role = _role
		}
		member.PhotoURL = user.PhotoURL
		member.Disabled = user.Disabled
		usersList = append(usersList, member)

	}

	c.JSON(http.StatusOK, usersList)
}

func (a AuthViews) UpdateUser(c *gin.Context) {

	member := models.Member{}
	uid := c.Param("uid")
	err := c.BindJSON(&member)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	client := models.GetFirebaseAuthClient(c)
	userToUpdate := (&auth.UserToUpdate{}).
		DisplayName(member.DisplayName).
		PhotoURL(member.PhotoURL).
		Disabled(member.Disabled).
		CustomClaims(map[string]interface{}{"role": member.Role})
	ur, err := member.UpdateUser(c.Request.Context(), client, uid, userToUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, ur)
}
