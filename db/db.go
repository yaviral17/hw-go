package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/yaviral17/hw-go/myLogs"
)

var db *sql.DB

func InitDB(dataSourceName string) error {
	myLogs.MyInfoLog("Connecting to database...")
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		myLogs.MyErrorLog(fmt.Sprintf("Error opening database: %q", err))
		return err
	}

	// Verify the connection
	err = db.Ping()
	if err != nil {
		myLogs.MyErrorLog(fmt.Sprintf("Error verifying connection with database: %q", err))
		return err
	}

	myLogs.MySuccessLog("Connected to database successfully 🚀")
	return nil
}

func GetDB() *sql.DB {
	return db
}
