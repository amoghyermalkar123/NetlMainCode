package db

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	mgo "gopkg.in/mgo.v2"
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
