package repository

import (
	_ "github.com/lib/pq"
)

// RepositoryPostgres db connectionn postgres
type RepositoryPostgres struct {
	Repository
}

// SetDbPostgres set db postgres and connection string
func SetDbPostgres(connectionString string) {
	setDb("postgres", connectionString)
}
