package postgres

import (
	//import the postgres driver
	_ "github.com/lib/pq"
)

type Driver struct {
	ScriptBuilder
}

func (Driver) GetDriverName() string {
	return "postgres"
}
