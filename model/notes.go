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
func GetNotes(tx *dbr.Tx, sectionID int, noteType string) []string {
	notes := []string{}
	cond := dbr.And(
		dbr.Eq("section_id", sectionID),
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
