package repository

import (
	_ "github.com/nakagami/firebirdsql"
)

// RepositoryFirebird db connectionn postgres
type RepositoryFirebird struct {
	Repository
}

// SetDbFirebird set db postgres and connection string
func SetDbFirebird(connectionString string) {
	setDb("firebirdsql", connectionString)
}
