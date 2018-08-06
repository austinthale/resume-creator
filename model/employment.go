package model

import "github.com/gocraft/dbr"

type Employment struct {
	Company      string   `json:"company"`
	DateAttended string   `json:"date_attended"`
	Position     string   `json:"position"`
	Notes        []string `json:"notes"`
	// Resume_ID	int 		`json:"resume_id"` ?????????????????
}

/****************************************************
* Save data from struct into PostgreSQL DB, and
* return err
****************************************************/
func (e *Employment) Save(tx *dbr.Tx) error {

	_, err := tx.InsertInto("Employment").
		Columns("company", "position", "date_attended", "resume_id").
		Record(e).
		Exec()

	return err
}
// TODO Find some way to get the notes for the Education object and append it to struct???
/****************************************************
* Load SQL data into PersonInfo struct based on the ID,
* and return ptr to that struct
****************************************************/
func (e *Employment) Load(tx *dbr.Tx, id int64) error {

	return tx.Select("company, date_attended, resume_id").
		From("Employment").
		Where("resume_id = ?", id).
		LoadOne(e)
}

type Employments []Employment


