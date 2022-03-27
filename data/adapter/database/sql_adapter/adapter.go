package sqladapter

import (
	"context"
	"database/sql"

	"github.com/mhogar/kiwi/data"
	"github.com/mhogar/kiwi/data/adapter/database"
)

type SQLDriver interface {
	ScriptBuilder

	// GetDriverName returns the name for the driver.
	GetDriverName() string
}

type SQLAdapter struct {
	database.DatabaseAdapter
	cancelFunc context.CancelFunc

	DB             *sql.DB
	SQLDriver      SQLDriver
	ContextFactory data.ContextFactory
}

// CreateSQLAdapter creates a new SQLAdapter with the provided driver.
func CreateSQLAdapter(driver SQLDriver) *SQLAdapter {
	adapter := &SQLAdapter{
		SQLDriver: driver,
	}
	adapter.Connection = adapter

	return adapter
}
