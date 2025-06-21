package db

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// New opens a database connection using the given DSN.
func New(dsn string) (*sql.DB, error) {
    return sql.Open("mysql", dsn)
}
