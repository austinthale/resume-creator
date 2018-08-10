package model

import (
	"github.com/gocraft/dbr"
	"github.com/sirupsen/logrus"
)

/**********************************************************
* Modular method, used to access details from database for
* any section type (Education, Employment, or Volunteer)
* for this Resume, and loads data into string array "notes".
***********************************************************/
func GetNotes(tx *dbr.Tx, sectionKey string, noteType string) []string {
	notes := []string{}
	cond := dbr.And(
		dbr.Eq("section_key", sectionKey),
		dbr.Eq("type", noteType),
	)
	_, err := tx.Select("detail").
		From("note").
		Where(cond).
		Load(&notes)
	if err != nil {
		logrus.Println("Error in GetNotes() at 'select' from database.")
		return nil
	}
	return notes
}
/**********************************************************
* Used to set the details into database for any section type
* whether (Education, Employment, or Volunteer) for this
* Resume, and stores data in corresponding table
***********************************************************/
func SetNote(tx *dbr.Tx, sectionKey string, noteType string, detail string) error {

	_, err := tx.InsertInto("note").
		Columns("section_key", "type", "detail").
		Values(sectionKey, noteType, detail).
		Exec()

	return err
}
/**********************************************************
* Removes all notes from database belonging to a section (whether Edu, Emp, or Vol)
***********************************************************/
func DeleteNotes(tx *dbr.Tx, sectionKey string, noteType string) error {
	cond := dbr.And(
		dbr.Eq("section_key", sectionKey),
		dbr.Eq("type", noteType),
	)
	// Delete all Notes for current section
	_, err := tx.DeleteFrom("note").Where(cond).Exec()
	return err
}