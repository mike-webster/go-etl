package models

import "context"

// Queryable is an interface used to move data from one database to another
type Queryable interface {
	// SourceQuery will return the sql that needs to be run against the source database.
	SourceQuery() string
	// DestinationSQL will return the sql that needs to be run to persist the information
	// in the Destination database.
	DestinationSQL(context.Context) []string
	// Errors will return all of the errors encountered during processing
	Errors() []string
	// Error will store an error
	Error(error)
}
