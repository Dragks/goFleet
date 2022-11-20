package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	// blank import for mysql driver
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	// connect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	// test db connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (dbAdapter Adapter) CloseConnection() {
	err := dbAdapter.db.Close()
	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}

func (dbAdapter Adapter) LogHistory(value float32, sensor string) error {
	queryString, args, err := sq.Insert("read_history").Columns("date", "value", "sensor").
		Values(time.Now(), value, sensor).ToSql()
	if err != nil {
		return err
	}

	_, err = dbAdapter.db.Exec(queryString, args...)
	if err != nil {
		return err
	}

	return nil
}
