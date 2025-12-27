package database

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/tbauriedel/resource-nexus-core/internal/config"
)

func TestGetDsn(t *testing.T) {
	expected := "postgresql://dummy:dummy@localhost:5432/test?sslmode=disable"

	c := config.Database{
		User:     "dummy",
		Password: "dummy",
		Address:  "localhost",
		Port:     5432,
		Name:     "test",
		TLSMode:  "disable",
	}

	dsn := getDsn(c)

	if dsn != expected {
		t.Fatalf("expected %s, got %s", expected, dsn)
	}
}

func TestNewDatabase(t *testing.T) {
	sqlOpen = func(driver, dsn string) (*sql.DB, error) {
		return nil, errors.New("boom")
	}

	_, err := NewDatabase(config.Database{}, nil)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestTestConnection(t *testing.T) {
	d, mock, _ := sqlmock.New(sqlmock.MonitorPingsOption(true))
	defer d.Close()

	mock.ExpectPing().WillReturnError(nil)

	db := SqlDatabase{
		database: d,
		logger:   nil,
	}

	err := db.TestConnection()
	if err != nil {
		t.Fatal(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal(err)
	}
}
