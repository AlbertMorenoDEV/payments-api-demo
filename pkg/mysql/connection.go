package mysql

import (
	"database/sql"
	"fmt"
	"time"
)

func Connect(user string, password string, host string, port uint16, databaseName string) (*sql.DB, error) {

	sqlAddress := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, host, port, databaseName)

	dbClient, err := sql.Open("mysql", sqlAddress)
	if err != nil {
		return nil, err
	}

	if err = dbClient.Ping(); err != nil {
		return nil, err
	}

	dbClient.SetConnMaxLifetime(3 * time.Minute)

	return dbClient, nil
}
