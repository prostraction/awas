package models

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Database struct {
	con *pgxpool.Pool
}

func (db *Database) Init() (err error) {
	// REPLACE !
	DB_URL := `postgres://postgres:postgres@localhost:5432/postgres`
	db.con, err = pgxpool.Connect(context.Background(), DB_URL)
	if err != nil {
		log.Fatal(err)
	}
	err = db.createTables()
	if err != nil {
		log.Fatal(err)
	}

	return
}

func (db *Database) noEmptyError(err error) error {
	if strings.ContainsAny(err.Error(), "no rows in result set") {
		return nil
	}
	return err
}

func (db *Database) createTables() (err error) {
	query := `CREATE TABLE IF NOT EXISTS number (id SERIAL, name VARCHAR(20), status text, type VARCHAR(8), id_type int);`
	err = db.con.QueryRow(context.Background(), query).Scan()
	if db.noEmptyError(err) != nil {
		return err
	}
	query = `CREATE TABLE IF NOT EXISTS number_history (id SERIAL, id_num int, status text, comment text);`
	err = db.con.QueryRow(context.Background(), query).Scan()
	return db.noEmptyError(err)
}

func (db *Database) GetNumberInfo(number int) (Number, error) {
	query := `SELECT id, name, status, type, id_type FROM number WHERE id = $1;`
	rows, err := db.con.Query(context.Background(), query, number)
	if err != nil {
		return Number{}, err
	}
	var num Number
	for rows.Next() {
		err := rows.Scan(&num.ID, &num.Name, &num.Status, &num.Type, &num.TypeID)
		if err != nil {
			log.Println(err.Error())
			return Number{}, err
		}
	}
	rows.Close()
	if num == (Number{}) {
		return Number{}, errors.New("no number found")
	}
	return num, nil
}

func (db *Database) GetNumberHistoryInfo(number int) (HistoryNumber, error) {
	query := `SELECT id, id_num, status, comment FROM number_history WHERE id = $1;`
	rows, err := db.con.Query(context.Background(), query, number)
	if err != nil {
		return HistoryNumber{}, err
	}
	var num HistoryNumber
	for rows.Next() {
		err = rows.Scan(&num.ID, &num.NumberID, &num.Status, &num.Comment)
		if err != nil {
			log.Println(err.Error())
			return HistoryNumber{}, err
		}
	}
	if num == (HistoryNumber{}) {
		err := errors.New("no history number found")
		return HistoryNumber{}, err
	}
	return num, nil
}
