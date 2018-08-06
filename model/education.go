package model

import "github.com/gocraft/dbr"

type Education struct {
	School       string   `json:"school"`
	DateAttended string   `json:"date_attended"`
	Notes        []string `json:"notes"`
	// Resume_ID	int 		`json:"resume_id"` ?????????????????
}

/****************************************************
* Save data from struct into PostgreSQL DB, and
* return err
****************************************************/
func (e *Education) Save(tx *dbr.Tx) error {

	_, err := tx.InsertInto("Education").
		Columns("school", "date_attended", "resume_id").
		Record(e).
		Exec()

	return err
}
// TODO Find some way to get the notes for the Education object and append it to struct???
/****************************************************
* Load SQL data into PersonInfo struct based on the ID,
* and return ptr to that struct
****************************************************/
func (e *Education) Load(tx *dbr.Tx, id int64) error {

	return tx.Select("school, date_attended, resume_id").
		From("Education").
		Where("resume_id = ?", id).
		LoadOne(e)
}

type Educations []Education
