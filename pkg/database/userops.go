package db

import (
	"context"
	model "netl/pkg/models/student"

	"gopkg.in/mgo.v2/bson"
)

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
