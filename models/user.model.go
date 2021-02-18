package models

import (
	"context"
	"os"

	"firebase.google.com/go/auth"
)

var (
	collectionName     string = os.Getenv("projectName")
	collectionNameUser        = collectionName + "UserCollection"
)

//User user struct
type User struct {
	Username   string
	Password   string
	Class      string
	CreateTime int64
	UpdateTime int64
}

type Member struct {
	UID          string      `json:"uid,omitempty"`
	DisplayName  string      `json:"displayName,omitempty"`
	Role         int         `json:"role,omitempty"`
	PhotoURL     string      `json:"photoUrl,omitempty"`
	Disabled     bool        `json:"disabled"`
	CustomClaims interface{} `json:"CustomClaims,omitempty"`
}

func (m Member) UpdateUser(ctx context.Context, client *auth.Client, uid string, userToUpdate *auth.UserToUpdate) (ur *auth.UserRecord, err error) {
	return client.UpdateUser(ctx, uid, userToUpdate)
}
