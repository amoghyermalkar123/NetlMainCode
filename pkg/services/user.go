package services

import (
	"context"
	"fmt"
	"math/rand"
	db "netl/pkg/database"
	model "netl/pkg/models/student"

	"time"

	"github.com/rs/zerolog"
	"gopkg.in/mgo.v2/bson"
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

func (us *UserService) startDatabaseSession() {
	// start a mongoDB session
	us.dtb.GetMongoDB()
}

// AuthorizeUser is a user service method to which takes in a payload of user information
func (us *UserService) AuthorizeUser(c context.Context, reg model.RegisterModel) error {
	// inits
	reg.UserId = generateUniqId()
	// starting a database session for this service method
	us.startDatabaseSession()
	// database ops
	dbErr := us.dtb.Dtb.C("users").Insert(reg)
	// checking if the performed database op had any errors
	if dbErr != nil {
		us.logger.Print("database problem occured")
		fmt.Printf("authorization error in user service, %+v\n", dbErr)
		return dbErr
	}
	// return nil incase of no errors
	return nil
}

// AuthenticateUser is a user service method that takes in a payload of login info and authenticates the user
func (us *UserService) AuthenticateUser(c context.Context, login model.LoginPayload) error {
	// inits
	userInfoModel := model.RegisterModel{}
	// starting a database session for this service method
	us.startDatabaseSession()
	// database query
	authenticateQuery := bson.M{
		"$and": []bson.M{
			{"email": &login.Email},
			{"password": &login.Password},
		},
	}
	// database ops
	dberror := us.dtb.Dtb.C("users").Find(authenticateQuery).One(&userInfoModel)
	// checking if the performed database op had any errors
	if dberror != nil {
		us.logger.Err(dberror)
		fmt.Printf("authentication error in user service, %+v\n", dberror)
		return dberror
	}
	// return nil incase of no errors
	return nil
}
