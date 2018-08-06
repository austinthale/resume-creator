package model

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

type Resumes []Resume