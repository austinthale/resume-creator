package model

import "github.com/gocraft/dbr"

type Education struct {
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


func (r *Resume) LoadEducations(tx *dbr.Tx, resumeID int64) error {
	educations := []Education{}
	_, err := tx.Select("*").
		From("education").
		Where("resume_id = ?", resumeID).
		Load(&educations)
	r.Educations = educations
	//r.Educations.GetEduNotes(tx, resumeID)
	return err
}


