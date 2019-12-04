package data

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Connection holds connection strings and databases
type Connection struct {
	// SourceDBDriverName is the driver to use to connect to the source database
	SourceDBDriverName string
	// SourceDB is a connection to the source database
	SourceDB *sqlx.DB
	// SourceDBString is the connection string to be used for the source database
	SourceDBConnectionString string
	// DestinationDBDriverName is the driver to use to connect to the destination database
	DestinationDBDriverName string
	// DestinationDB is a connection to the destination database
	DestinationDB *sqlx.DB
	// DestinationDBConnectionString is the connection string to be used for the destination database
	DestinationDBConnectionString string
}

// Initialize will open database connections with the provided connection strings
func (db *Connection) Initialize() *[]error {
	ret := []error{}
	db1, err := sqlx.Open(db.SourceDBDriverName, db.SourceDBConnectionString)
	db.SourceDB = db1
	if err != nil {
		ret = append(ret, err)
	}

	db2, err := sqlx.Open(db.DestinationDBDriverName, db.DestinationDBConnectionString)
	db.DestinationDB = db2
	if err != nil {
		ret = append(ret, err)
	}

	if len(ret) > 0 {
		return &ret
	}
	return nil
}

// SourceSelect will run a query against the source database and return the results
func (db *Connection) SourceSelect(query string) (*sqlx.Rows, error) {
	res, err := db.SourceDB.Queryx(query)
	if err != nil {
		fmt.Println("--- Error query: ", query)
	}
	return res, err
}

// DestinationInsert will run a query against the destination database and return the results
func (db *Connection) DestinationInsert(query string) (interface{}, error) {
	res, err := db.DestinationDB.Exec(query)
	return res, err
}

// DestinationIDSelect will run a query against the destination database and store the
// results in the given interface, an error is returned if one is encountered.
func (db *Connection) DestinationIDSelect(query string) (*[]int64, error) {
	var ret []int64
	err := db.DestinationDB.Select(&ret, query)
	return &ret, err
}
