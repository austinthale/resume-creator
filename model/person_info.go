package model

import "github.com/gocraft/dbr"

type PersonInfo struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
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

	return tx.Select("name, address, phone, email").
		From("Resume").
		Where("id = ?", id).
		LoadOne(p)
}