package model

import (
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"net/http"
)

type Resume struct {
	PersonInfo  PersonInfo   `json:"person_info"`
	Educations  []Education  `json:"educations"`
	Employments []Employment `json:"employments"`
	Volunteers  []Volunteer  `json:"volunteers"`
}

func NewResume(info PersonInfo, educations []Education, employments []Employment, volunteers []Volunteer) *Resume {
	return &Resume{
		PersonInfo:    info,
		Educations:    educations,
		Employments:   employments,
		Volunteers:	   volunteers,
	}
}

func (r *Resume) Load(tx *dbr.Tx, resumeID int64) error {
	error := false
	if err := r.PersonInfo.GetPersonInfo(tx, resumeID); err != nil { error = true }
	if err := r.GetEducations(tx, resumeID); err != nil { error = true }
	if err := r.GetEmployments(tx, resumeID); err != nil { error = true }
	if err := r.GetVolunteers(tx, resumeID); err != nil { error = true }
	if error != false {
		return echo.NewHTTPError(http.StatusNotFound, "Error loading data from database.")
	}
	return nil
}

func (r *Resume) Store(tx *dbr.Tx, resumeID int64) error {
	error := false
	if err := r.PersonInfo.SetPersonInfo(tx, resumeID); err != nil { error = true }
	if err := r.SetEducations(tx, resumeID); err != nil { error = true }
	if err := r.SetEmployments(tx, resumeID); err != nil { error = true }
	if err := r.SetVolunteers(tx, resumeID); err != nil { error = true }
	if error != false {
		return echo.NewHTTPError(http.StatusNotFound, "Error saving data in database.")
	}
	return nil
}




