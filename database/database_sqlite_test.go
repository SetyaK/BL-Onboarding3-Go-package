package database

import (
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestNewSessionSqlite3(t *testing.T) {
	os.Setenv("DATABASE_ADAPTER", "sqlite3")

	sess, err := NewSession()
	if err != nil {
		t.Fatal("Cannot estabilish session. " + err.Error())
	}

	err = sess.Ping()
	if err != nil {
		t.Fatal("Ping session failed. " + err.Error())
	}

	//Schema Migration Test
	m := Migration{sess}
	_, err = m.Migrate()
	if err != nil {
		t.Fatal("Migration failed." + err.Error())
	}
}
