package services

import (
	"context"
	"math/rand"
	db "netl/pkg/database"
	model "netl/pkg/models/student"

	"time"

	"github.com/rs/zerolog"
)

// UserService type embeds the database type to create sessions and database queries
type UserService struct {
	dtb    db.DB
	logger zerolog.Logger
}

func generateId(min, max int) int {
	return min + rand.Intn(max-min)
}

func generateUniqId() int {
	rand.Seed(time.Now().UnixNano())
	uid := generateId(0, 2000)

	return uid
}

// AuthorizeUser is a user service method to which takes in a payload of user information
func (us *UserService) AuthorizeUser(c context.Context, reg model.RegisterModel) error {
	// generating a unique random user id
	reg.UserId = generateUniqId()
	// database ops
	dbErr := us.dtb.InsertUser(c, reg)
	// checking if the performed database op had any errors
	if dbErr != nil {
		us.logger.Debug().Err(dbErr).Msg("user service failed to authorize")
		return dbErr
	}
	// return nil incase of no errors
	return nil
}

// AuthenticateUser is a user service method that takes in a payload of login info and authenticates the user
func (us *UserService) AuthenticateUser(c context.Context, login model.LoginPayload) error {
	// database op
	dberror := us.dtb.FindUser(c, login)
	// checking if the performed database op had any errors
	if dberror != nil {
		us.logger.Debug().Err(dberror).Msg("user service failed to authenticate")
		return dberror
	}
	// return nil incase of no errors
	return nil
}
