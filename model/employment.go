package model

import "github.com/gocraft/dbr"

type Employment struct {
	ID 			 int	  `db:"id"`
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
func (e *Employment) Save(tx *dbr.Tx) error {

	_, err := tx.InsertInto("Employment").
		Columns("company", "position", "date_attended", "resume_id").
		Record(e).
		Exec()

	return err
}

/****************************************************
* Retrieves each Employment section from database for
* this Resume, and loads data into struct (JSON format).
****************************************************/
func (r *Resume) GetEmployments(tx *dbr.Tx, resumeID int64) error {
	employments := []Employment{}
	_, err := tx.Select("*").
		From("employment").
		Where("resume_id = ?", resumeID).
		Load(&employments)
	r.Employments = employments

	for idx, emp := range r.Employments {
		r.Employments[idx].Notes = GetNotes(tx, emp.ID, "employment")
	}
	return err
}

