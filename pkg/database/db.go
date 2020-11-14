package db

import (
	"context"
	"fmt"
	model "netl/pkg/models/student"
	"os"

	"github.com/rs/zerolog"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DB type embeds its own local logger and gives required services methods that perform database operations
type DB struct {
	logger zerolog.Logger
}

// database session variable
var db *mgo.Database

// initiliaze a variable for database sessions at the start of the server
func init() {
	host := "localhost"
	dbName := "netldb"

	session, err := mgo.Dial(host)
	if err != nil {
		fmt.Println("session err:", err)
		os.Exit(2)
	}
	db = session.DB(dbName)
}

// InsertUser takes in a user model and persists it into the given collection in the database
func (dba *DB) InsertUser(c context.Context, reg model.RegisterModel) error {
	// database insert operation
	dbInsertOp := db.C("users").Insert(reg)
	// check if there's any error while persisting
	if dbInsertOp != nil {
		dba.logger.Debug().Err(dbInsertOp).Msg("authorzation failed")
		return dbInsertOp
	}
	// return nil indicating operation successfull
	return nil
}

func (dba *DB) FindUser(c context.Context, login model.LoginPayload) error {
	// creating the model instance
	userInfoModel := model.RegisterModel{}
	// database authentication query
	authenticateQuery := bson.M{
		"$and": []bson.M{
			{"email": &login.Email},
			{"password": &login.Password},
		},
	}
	// performing the database operation
	dbFindOP := db.C("users").Find(authenticateQuery).One(&userInfoModel)
	// check if there's any error while retrieving
	if dbFindOP != nil {
		dba.logger.Debug().Err(dbFindOP).Msg("authentication failed")
		return dbFindOP
	}
	// return nil indicating operation successfull
	return nil
}
