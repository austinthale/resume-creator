package model

import (
	"github.com/gocraft/dbr"
	"github.com/austinthale/resume-creator/secure"
	"net/http"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type User struct {
	Username   	string 	`json:"username" db:"username"`
	Password	string  `json:"password" db:"password"`
}

func (u *User) Validate(tx *dbr.Tx) error {
	// Make new user obj to hold info from database, then compare to bound object we passed in (u)
	DBInfo := new(User)

	err := tx.Select("password").
		From("public.user").
		Where("username = ?", u.Username).
		LoadOne(DBInfo)
	if err != nil {
		logrus.Println("Error retrieving user credentials from DB")
		return err
	}

	if err != nil {
		logrus.Println("Error hashing password: ")
		return err
	}
	if secure.CheckPasswordHash(u.Password, DBInfo.Password) {
		return nil
	} else {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
	}
}
