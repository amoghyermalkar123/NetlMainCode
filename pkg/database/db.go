package db

import (
	"fmt"
	"os"

	mgo "gopkg.in/mgo.v2"
)

type DB struct {
	Dtb *mgo.Database
}

var db *mgo.Database

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

// GetMongoDB function to return DB connection
func (dab *DB) GetMongoDB() {
	dab.Dtb = db
}
