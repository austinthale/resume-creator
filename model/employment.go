package model

import (
	"github.com/gocraft/dbr"
	"github.com/austinthale/resume-creator/pkg"
)

type Employment struct {
	ID           int      `db:"id"`
	EmpKey       string   `json:"employment_key" db:"employment_key"`
	Company      string   `json:"company" db:"company"`
	DateAttended string   `json:"date_attended" db:"date_attended"`
	Position     string   `json:"position" db:"position"`
	Notes        []string `json:"notes"`
	ResumeID     int      `json:"resume_id" db:"resume_id"`
	IsDeleted    bool     `json:"is_deleted"`
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
		r.Employments[idx].Notes = GetNotes(tx, emp.EmpKey, "employment")
	}
	return err
}

/****************************************************
* Set Employments data from struct into PostgreSQL DB, and
* return err
****************************************************/
func (r *Resume) SetEmployments(tx *dbr.Tx, id int64) error {
	// Loop through employments to Delete, Insert, or Update
	for _, emp := range r.Employments {
		if emp.IsDeleted { // If IsDeleted is true, remove from DB
			DeleteEmployment(tx, emp.EmpKey)
		} else if emp.EmpKey != "" { // If this employment HAS an ID already, update it
			UpdateEmployment(tx, emp)
		} else if emp.EmpKey == "" {					// Else, it's null. Insert new Employment
			// Generate random key for this employment section
			emp.EmpKey = pkg.RandSeq(16)
			emp.ResumeID = int(id)
			InsertEmployment(tx, emp)
		}
		// Delete all Notes for current section of employment
		if err := DeleteNotes(tx, emp.EmpKey, "employment"); err != nil {
			return err
		}
		if !(emp.IsDeleted){
			// Insert notes for current section of employment (NOTE: ID's for employments have been updated!
			for _, note := range emp.Notes {
				if err := SetNote(tx, emp.EmpKey, "employment", note); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
/******************************************************
* Removes Employment data from database
******************************************************/
func DeleteEmployment(tx *dbr.Tx, empKey string) error {
	_, err := tx.DeleteFrom("employment").
		Where("employment_key = ?", empKey).
		Exec()
	return err
}
/*******************************************************
* Updates already existing Employment data from database
*******************************************************/
func UpdateEmployment(tx *dbr.Tx, emp Employment) error {
	_, err := tx.Update("employment").
		Set("company", emp.Company).
		Set("position", emp.Position).
		Set("date_attended", emp.DateAttended).
		Where("employment_key = ?", emp.EmpKey).
		Exec()
	return err
}
/*******************************************************
* Inserts non-existing Employment data into database
*******************************************************/
func InsertEmployment(tx *dbr.Tx, emp Employment) error {
	_, err := tx.InsertInto("employment").
		Columns("company", "position", "date_attended", "resume_id", "employment_key").
		Record(&emp).
		Exec()
	return err
}