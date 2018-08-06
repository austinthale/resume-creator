package db

import (
	"os"
	"github.com/gocraft/dbr"
	"github.com/sirupsen/logrus"
)

func Init() *dbr.Session {

	session := getSession()

	return session
}

func getSession () *dbr.Session {
	dbUrl := os.Getenv("DATABASE_URL")
	db, err := dbr.Open("postgres", dbUrl, nil)
	//checkErr(err)
	if err != nil {
		logrus.Error(err)
	} else {
		session := db.NewSession(nil)
		return session
	}
	/*
	//defer db.Close()
	*/
	return nil
}