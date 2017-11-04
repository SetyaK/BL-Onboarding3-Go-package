package database

import (
	"os"

	"github.com/gocraft/dbr"
)

// NewSession instantiates a Session for the Connection
func NewSession() (*dbr.Session, error) {
	driver := os.Getenv("DATABASE_ADAPTER")
	port := os.Getenv("DATABASE_PORT")
	dataSourceName := ""
	switch driver {
	case "mysql":
		if port == "" {
			port = "3306"
		}
		dataSourceName = os.Getenv("DATABASE_USERNAME") + ":" + os.Getenv("DATABASE_PASSWORD") + "@(" + os.Getenv("DATABASE_HOST") + ":" + (string(port)) + ")/" + os.Getenv("DATABASE_NAME")
	case "sqlite3":
		dataSourceName = ":memory:"
	}
	conn, err := dbr.Open(driver, dataSourceName, nil)
	if err != nil {
		return nil, err
	}
	session := conn.NewSession(nil)
	return session, nil
}
