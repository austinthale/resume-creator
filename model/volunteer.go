package model

import (
	"github.com/gocraft/dbr"
)

type Volunteer struct {
	ID           int      `db:"id"`
	VolKey       string   `json:"volunteer_key" db:"volunteer_key"`
	Company      string   `json:"company" db:"company"`
	DateAttended string   `json:"date_attended" db:"date_attended"`
	Position     string   `json:"position" db:"position"`
	Notes        []string `json:"notes"`
	ResumeID     int      `json:"resume_id" db:"resume_id"`
	IsDeleted    bool     `json:"is_deleted"`
}

/****************************************************
* Retrieves each Volunteer section from database for
* this Resume, and loads data into struct (JSON format).
****************************************************/
func (r *Resume) GetVolunteers(tx *dbr.Tx, resumeID int64) error {
	volunteers := []Volunteer{}
	_, err := tx.Select("*").
		From("volunteer").
		Where("resume_id = ?", resumeID).
		Load(&volunteers)
	r.Volunteers = volunteers

	for idx, vol := range r.Volunteers {
		r.Volunteers[idx].Notes = GetNotes(tx, vol.VolKey, "volunteer")
	}
	return err
}

/****************************************************
* Set Volunteers data from struct into PostgreSQL DB, and
* return err
****************************************************/
func (r *Resume) SetVolunteers(tx *dbr.Tx, id int64) error {
	// Loop through volunteers to Delete, Insert, or Update
	for _, vol := range r.Volunteers {
		if vol.IsDeleted { // If IsDeleted is true, remove from DB
			DeleteVolunteer(tx, vol.VolKey)
		} else if vol.VolKey != "" { // If this volunteer HAS an ID already, update it
			UpdateVolunteer(tx, vol)
		} else if vol.VolKey == "" { // Else, it's null. Insert new Volunteer
			// Generate random key for this education section
			vol.VolKey = RandSeq(16)
			vol.ResumeID = int(id)
			InsertVolunteer(tx, vol)
		}
		// Delete all Notes for current section of volunteer
		if err := DeleteNotes(tx, vol.VolKey, "volunteer"); err != nil { // before I used 'i + 1' instead of vol.ID
			return err
		}
		if !(vol.IsDeleted) {
			// Insert notes for current section of volunteer (NOTE: ID's for volunteers have been updated!
			for _, note := range vol.Notes {
				if err := SetNote(tx, vol.VolKey, "volunteer", note); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
/******************************************************
* Removes Volunteer data from database
******************************************************/
func DeleteVolunteer(tx *dbr.Tx, volKey string) error {
	_, err := tx.DeleteFrom("volunteer").
		Where("volunteer_key = ?", volKey).
		Exec()
	return err
}
/*******************************************************
* Updates already existing Volunteer data from database
*******************************************************/
func UpdateVolunteer(tx *dbr.Tx, vol Volunteer) error {
	_, err := tx.Update("volunteer").
		Set("company", vol.Company).
		Set("position", vol.Position).
		Set("date_attended", vol.DateAttended).
		Where("volunteer_key = ?", vol.VolKey).
		Exec()
	return err
}
/*******************************************************
* Inserts non-existing Volunteer data into database
*******************************************************/
func InsertVolunteer(tx *dbr.Tx, vol Volunteer) error {
	_, err := tx.InsertInto("volunteer").
		Columns("company", "position", "date_attended", "resume_id", "volunteer_key").
		Record(&vol).
		Exec()
	return err
}