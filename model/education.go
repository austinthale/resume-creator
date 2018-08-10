package model

import (
	"github.com/gocraft/dbr"
)
// The EduKey is a work-around because postgres couldn't get the ID's for the section
// in order to insert notes (would've had to use INSERT with RETURNING, to get last ID)
type Education struct {
	ID           int      `db:"id"`
	EduKey       string   `json:"education_key" db:"education_key"`
	School       string   `json:"school" db:"school"`
	DateAttended string   `json:"date_attended" db:"date_attended"`
	Notes        []string `json:"notes"`
	ResumeID     int      `json:"resume_id" db:"resume_id"`
	IsDeleted    bool     `json:"is_deleted"`
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
		r.Educations[idx].Notes = GetNotes(tx, edu.EduKey, "education")
	}
	return err
}

/****************************************************
* Set Educations data from struct into PostgreSQL DB, and
* return err
****************************************************/
func (r *Resume) SetEducations(tx *dbr.Tx, id int64) error {
	// Loop through educations to Delete, Insert, or Update
	for _, edu := range r.Educations {
		if edu.IsDeleted { // If IsDeleted is true, remove from DB
			DeleteEducation(tx, edu.EduKey)
		} else if edu.EduKey != "" { // If this education HAS an ID already, update it
			UpdateEducation(tx, edu)
		} else if edu.EduKey == "" {					// Else, it's null. Insert new Education
			// Generate random key for this education section
			edu.EduKey = RandSeq(16)
			edu.ResumeID = int(id)
			InsertEducation(tx, edu)
		}
		// Delete all Notes for current section of education
		if err := DeleteNotes(tx, edu.EduKey, "education"); err != nil {
			return err
		}
		if !(edu.IsDeleted){
			// Insert notes for current section of education (NOTE: ID's for educations have been updated!
			for _, note := range edu.Notes {
				if err := SetNote(tx, edu.EduKey, "education", note); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
/******************************************************
* Removes Education data from database
******************************************************/
func DeleteEducation(tx *dbr.Tx, eduKey string) error {
	_, err := tx.DeleteFrom("education").
		Where("education_key = ?", eduKey).
		Exec()
	return err
}
/*******************************************************
* Updates already existing Education data from database
*******************************************************/
func UpdateEducation(tx *dbr.Tx, edu Education) error {
	_, err := tx.Update("education").
		Set("school", edu.School).
		Set("date_attended", edu.DateAttended).
		Where("education_key = ?", edu.EduKey).
		Exec()
	return err
}
/*******************************************************
* Inserts non-existing Education data into database
*******************************************************/
func InsertEducation(tx *dbr.Tx, edu Education) error {
	_, err := tx.InsertInto("education").
		Columns("school", "date_attended", "resume_id", "education_key").
		Record(&edu).
		Exec()
	return err
}