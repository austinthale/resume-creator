package model

import (
	"github.com/gocraft/dbr"
)

type PersonInfo struct {
	Name    string `json:"name" db:"name"`
	Address string `json:"address" db:"address"`
	Phone   string `json:"phone" db:"phone"`
	Email   string `json:"email" db:"email"`
}

/****************************************************
* Save data from struct into PostgreSQL DB, and
* return err
****************************************************/
func (p *PersonInfo) Save(tx *dbr.Tx) error {

	_, err := tx.InsertInto("Resume").
		Columns("name", "address", "phone", "email").
		Record(p).
		Exec()

	return err
}

/****************************************************
* Load SQL data into PersonInfo struct based on the ID,
* and return ptr to that struct
****************************************************/
func (p *PersonInfo) Load(tx *dbr.Tx, id int64) error {

	return tx.Select("name", "address", "phone", "email").
		From("resume").
		Where("id = ?", id).
		LoadOne(p)
}