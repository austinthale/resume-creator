package model

import (
	"github.com/gocraft/dbr"
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
	r.LoadEducations(tx, resumeID)
	r.LoadEmployments(tx, resumeID)
	r.LoadVolunteers(tx, resumeID)
	//TODO r.GetEmpNotes(), r.GetEduNotes(), r.GetVolNotes()
	return nil
}




