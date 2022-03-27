package sqladapter

import (
	"context"
	"database/sql"

	"github.com/mhogar/kiwi/common"
)

// OpenConnection opens the connection to SQL database server using the fields from the database config.
// Initializes the adapter's context and cancel function, as well as its db instance.
// Returns any errors.
func (a *SQLAdapter) OpenConnection() error {
	connectionStr := "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"

	a.ContextFactory.Context, a.cancelFunc = context.WithCancel(context.Background())
	a.ContextFactory.Timeout = 5000

	//connect to the db
	db, err := sql.Open(a.SQLDriver.GetDriverName(), connectionStr)
	if err != nil {
		return common.ChainError("error opening database connection", err)
	}

	a.DB = db

	return nil
}

// CloseConnection closes the connection to the SQL database server and resets its db instance.
// The adapter also calls its cancel function to cancel any child requests that may still be running.
// Neither the adapter's db instance or context should be used after calling this function.
// Returns any errors.
func (a *SQLAdapter) CloseConnection() error {
	err := a.DB.Close()
	if err != nil {
		return common.ChainError("error closing database connection", err)
	}

	//cancel any remaining requests that may still be running
	a.cancelFunc()

	//clean up resources
	a.DB = nil

	return nil
}

// Ping pings the SQL database server to verify it can still be reached.
// Returns an error if it cannot, or if any other errors are encountered.
func (DB *SQLAdapter) Ping() error {
	ctx, cancel := DB.ContextFactory.CreateStandardTimeoutContext()
	err := DB.DB.PingContext(ctx)
	cancel()

	if err != nil {
		return common.ChainError("error pinging database", err)
	}

	return nil
}
