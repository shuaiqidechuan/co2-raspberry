package mysql

import (
	"database/sql"
	"errors"
)

const (
	mysqlDioxideInsert = iota
)

var (
	errInvalidInsert = errors.New("errInvalidInsert")
)

var (
	dioxideSQLString = []string{
		// `INSERT INTO test (id) VALUES (?)`,
		"INSERT INTO Co2_test (dioxide) VALUES (?)",
	}
)

func InsertDioxide(db *sql.DB, dioxideDensity int) error {
	result, err := db.Exec(dioxideSQLString[mysqlDioxideInsert], dioxideDensity)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errInvalidInsert
	}

	return nil
}