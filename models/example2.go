package models

import (
	"fmt"
	"context"
)

// Example2 is an example representing a class that we might want to move from
// one database to another.
type Example2 struct {
	ID int64 `db:"id"`
	Name string `db:"name"`
	errors []error
}

// Begin Queryable Methods ---- ---- ----
//

// SourceQuery will return the sql that needs to be run against the source database.
func (e *Example2) SourceQuery() string {
	return "SELECT id, name FROM example2;"
}

// DestinationSQL will return the sql that needs to be run to persist the information
// in the Destination database.
// Note: This will return a slice of sql statements that need to be executed
func (e *Example2) DestinationSQL(context.Context) []string {
	return []string{
		fmt.Sprintf("INSERT INTO example2 (source_id, name) VALUES (%v, '%v-transformed');", e.ID, e.Name),
	}
}

// Errors will return all of the errors encountered during processing
func (e *Example2) Errors() []string {
	var ret []string
	for _, er := range e.errors {
		ret = append(ret, er.Error())
	}
	return ret
}

// Error will store an error
func (e *Example2) Error(er error) {
	e.errors = append(e.errors, er)
}

//
// End Queryable Methods ---- ---- ----