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
* SetPersonInfo data from struct into PostgreSQL DB, and
* return err
****************************************************/
func (p *PersonInfo) SetPersonInfo(tx *dbr.Tx, id int64) error {

	_, err := tx.Update("resume").
		Set("name", p.Name).
		Set("address", p.Address).
		Set("phone", p.Phone).
		Set("email", p.Email).
		Where("id = ?", id).
		Exec()

	return err
}

/****************************************************
* GetPersonInfo SQL data into PersonInfo struct based on the ID,
* and return ptr to that struct
****************************************************/
func (p *PersonInfo) GetPersonInfo(tx *dbr.Tx, id int64) error {

	return tx.Select("name", "address", "phone", "email").
		From("resume").
		Where("id = ?", id).
		LoadOne(p)
}