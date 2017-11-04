package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
)

var dropall = "DROP TABLE IF EXISTS products"

var products = `CREATE TABLE products (
		product_id %s,
		name varchar(255) NOT NULL,
		description varchar(255) NULL,
		stock integer DEFAULT 0
	)`

// Migration for database migration
type Migration struct {
	Session *dbr.Session
}

// Migrate create the database schema
func (m Migration) Migrate() (sql.Result, error) {
	if os.Getenv("ENV") == "development" {
		_, err := m.Session.Exec(dropall)
		if err != nil {
			return nil, err
		}
	}

	var autoIncrementType string
	switch m.Session.Dialect {
	case dialect.MySQL:
		autoIncrementType = "serial PRIMARY KEY"
	case dialect.SQLite3:
		autoIncrementType = "integer PRIMARY KEY"
	}

	productsQuery := fmt.Sprintf(products, autoIncrementType)
	return m.Session.Exec(productsQuery)
}
