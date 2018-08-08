package model

import (
	"github.com/gocraft/dbr"
)

type Education struct {
	ID		 	 int	  `db:"id"`
	School       string   `json:"school" db:"school"`
	DateAttended string   `json:"date_attended" db:"date_attended"`
	Notes        []string `json:"notes"`
	ResumeID	 int 	  `json:"resume_id" db:"resume_id"`
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

/****************************************************
* Retrieves each Education section from database for
* this Resume, and loads data into struct (JSON format).
****************************************************/
func (r *Resume) GetEducations(tx *dbr.Tx, resumeID int64) error {
	educations := []Education{}
	_, err := tx.Select("*").
		From("education").
		Where("resume_id = ?", resumeID).
		Load(&educations)
	r.Educations = educations

	for idx, edu := range r.Educations {
		r.Educations[idx].Notes = GetNotes(tx, edu.ID, "education")
	}

	return err
}


