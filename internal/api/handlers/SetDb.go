package handlers

import "database/sql"

// DB is the package‐level *sql.DB used by your handlers.
var DB *sql.DB

// SetDB initializes the package‐level DB. Call this once in main().
func SetDB(conn *sql.DB) {
	DB = conn
}
