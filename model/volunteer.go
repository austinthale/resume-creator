package model

import "github.com/gocraft/dbr"

type Volunteer struct {
	Company      string   `json:"company" db:"company"`
	DateAttended string   `json:"date_attended" db:"date_attended"`
	Position     string   `json:"position" db:"position"`
	Notes        []string `json:"notes"`
	ResumeID	 int 	  `json:"resume_id" db:"resume_id"`
}

/****************************************************
* Save data from struct into PostgreSQL DB, and
* return err
****************************************************/
func (v *Volunteer) Save(tx *dbr.Tx) error {

	_, err := tx.InsertInto("Volunteer").
		Columns("company", "position", "date_attended", "resume_id").
		Record(v).
		Exec()

	return err
}
// TODO Find some way to get the notes for the Education object and append it to struct???
/****************************************************
* Load SQL data into PersonInfo struct based on the ID,
* and return ptr to that struct
****************************************************/
func (v *Volunteer) Load(tx *dbr.Tx, id int64) error {

	return tx.Select("company, date_attended, resume_id").
		From("Employment").
		Where("resume_id = ?", id).
		LoadOne(v)
}

func (r *Resume) LoadVolunteers(tx *dbr.Tx, resumeID int64) error {
	volunteers := []Volunteer{}
	_, err := tx.Select("*").
		From("volunteer").
		Where("resume_id = ?", resumeID).
		Load(&volunteers)
	r.Volunteers = volunteers
	return err
}