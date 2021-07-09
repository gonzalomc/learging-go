package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Obtener conexión a base de datos
// !TODO: Trabajar con variables de entorno
func GetConnection() (*sql.DB, error) {
	uri := "postgres://postgres:d3m0g0@127.0.0.1:5432/survey?sslmode=disable"
	return sql.Open("postgres", uri)
}
