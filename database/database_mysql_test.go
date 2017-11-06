package database

import (
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestNewSessionMysql(t *testing.T) {
	os.Setenv("DATABASE_ADAPTER", "mysql")
	os.Setenv("DATABASE_HOST", "localhost")
	os.Setenv("DATABASE_USERNAME", "root")
	os.Setenv("DATABASE_PASSWORD", "root")
	os.Setenv("DATABASE_NAME", "sys")
	os.Setenv("ENV", "development")

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

func TestInvalidSessionMysql(t *testing.T) {
	os.Setenv("DATABASE_ADAPTER", "")

	_, err := NewSession()
	if err == nil {
		t.Fatal("Session should fail")
	}
}
