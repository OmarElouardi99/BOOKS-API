package driver

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	SQL *sqlx.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifeTime = 5 * time.Minute

// creates db pool
func ConnectSQL(dsn string) (*DB, error) {

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifeTime)
	db.SetMaxOpenConns(maxOpenDbConn)

	err = testDB(db)
	if err != nil {
		return nil, err
	}

	dbConn.SQL = db
	return dbConn, nil
}

// tries to ping db
func testDB(d *sqlx.DB) error {

	err := d.Ping()
	if err != nil {
		fmt.Println("ERROR !", err)
		return err
	}

	fmt.Println("Database pinged successfully! ***")
	return nil
}
