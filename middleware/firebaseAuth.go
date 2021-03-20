package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

func FirebaseAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		authString := c.GetHeader("Authorization") // ["Authorization"][0]
		kv := strings.Split(authString, " ")
		if len(kv) != 2 || kv[0] != "Bearer" {
			statusCode := http.StatusUnauthorized
			c.JSON(statusCode, http.StatusText(statusCode))
			c.Abort()
			return
		}
		tokenString := kv[1]

		// Verify the session cookie. In this case an additional check is added to detect
		// if the user's Firebase session was revoked, user deleted/disabled, etc.

		client, err := newClient(c.Request.Context())
		if err != nil {
			log.Println(err)
			statusCode := http.StatusUnauthorized
			c.JSON(statusCode, http.StatusText(statusCode))
			c.Abort()
			return
		}

		decoded, err := client.VerifyIDTokenAndCheckRevoked(c.Request.Context(), tokenString)
		if err != nil {
			//idtoken is unavailable. Force user to login.
			log.Println(err)
			statusCode := http.StatusUnauthorized
			c.JSON(statusCode, http.StatusText(statusCode))
			c.Abort()
			return
		}

		c.Set("token", decoded)
		c.Set("role", decoded.Claims["role"])
		c.Set("username", decoded.Claims["name"])
		c.Next()
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Get("role")
		if !ok {
			statusCode := http.StatusForbidden
			c.JSON(statusCode, "no role in context")
			c.Abort()
			return
		}

		_role, err := strconv.Atoi(fmt.Sprintf("%v", role))
		if err != nil {
			statusCode := http.StatusForbidden
			c.JSON(statusCode, err)
			c.Abort()
			return
		}

		if !(_role == 1 || _role == 2) {
			statusCode := http.StatusForbidden
			c.JSON(statusCode, "Forbidden")
			c.Abort()
			return
		}
		c.Next()
	}
}

func newClient(ctx context.Context) (*auth.Client, error) {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, err
	}
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	return client, err
}
