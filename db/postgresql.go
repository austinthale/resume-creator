package db

import (
	"github.com/gocraft/dbr"
	"github.com/sirupsen/logrus"
	"fmt"
	"github.com/austinthale/resume-creator/conf"
)

func Init() *dbr.Session {

	session := getSession()

	return session
}

func getSession () *dbr.Session {
	dbinfo := fmt.Sprintf("dbname=%s host=%s port=%s user=%s password=%s sslmode=%s",
		conf.DB_NAME, conf.DB_HOST, conf.DB_PORT, conf.DB_USER, conf.DB_PASSWORD,  conf.SSL_MODE)
	db, err := dbr.Open("postgres", dbinfo, nil)
	// Could also use:
	//dbUrl := os.Getenv("DATABASE_URL")
	//db, err := dbr.Open("postgres", dbUrl, nil)
	//checkErr(err)
	if err != nil {
		logrus.Println("DB session FAILED")
		logrus.Error(err)
	} else {
		//defer db.Close()
		logrus.Println("DB session SUCCESS")
		session := db.NewSession(nil)
		return session
	}

	return nil
}