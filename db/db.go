package db

import (
	"database/sql"
	"fmt"

	"github.com/rafaeljusto/cctldcentral/config"
)

// Connection database connection.
var Connection *sql.DB

// Connect performs the database connection.
func Connect() (err error) {
	connParams := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		config.CCTLDCentral.Database.Username,
		config.CCTLDCentral.Database.Password,
		config.CCTLDCentral.Database.Host,
		config.CCTLDCentral.Database.Name,
	)

	Connection, err = sql.Open("postgres", connParams)
	return
}
