package testutils

import (
	"database/sql"
	"os"
	"time"
)

// NewDB create and returns new *sql.DB from environment
// in your tests you need to add sql drivers manually
// _ "github.com/ClickHouse/clickhouse-go"
// _ "github.com/go-sql-driver/mysql"
// _ "github.com/lib/pq"
func NewSqlDB(t Tester, driverName, envName string) *sql.DB {
	dsn := os.Getenv(envName)
	if dsn == "" {
		t.Skip("empty dsn")
	}

	db, err := sql.Open(driverName, dsn)
	if err != nil {
		t.Fatalf("opening sql connection failed: %s", err)
	}

	db.SetConnMaxLifetime(time.Second)

	return db
}
