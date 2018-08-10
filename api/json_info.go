package api

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/austinthale/resume-creator/model"
	"github.com/gocraft/dbr"
	"strconv"
	"github.com/sirupsen/logrus"
)
// TODO build the r variable by accessing database and convert to structs

func GetResumeInfo() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		stringID := c.Param("id")
		id, err := strconv.ParseInt(stringID, 10, 64)
		if err != nil {
			return err
		}

		var r = model.Resume{}
		tx := c.Get("Tx").(*dbr.Tx)

		// Get all the data into resume from DB
		r.Load(tx, id)

		return c.JSON(http.StatusOK, r)
	}
}

func SetResumeInfo() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		stringID := c.Param("id")
		id, err := strconv.ParseInt(stringID, 10, 64)
		if err != nil {
			return err
		}
		tx := c.Get("Tx").(*dbr.Tx)
		res := new(model.Resume)
		// Bind gets the JSON data from the page and binds it to the Resume object
		if err = c.Bind(res); err != nil {
			logrus.Println("Error binding data")
			return
		}
		r := *res	// 'r' now reflects the input that we used BIND to get, we can update the database

		// Store the resume data in DB
		r.Store(tx, id)

		return c.JSON(http.StatusOK, r)
	}
}

/*func DisplayJSON() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		stringID := c.Param("id")
		id, err := strconv.ParseInt(stringID, 10, 64)
		if err != nil {
			return err
		}

		var r = model.Resume{}

		// personal info
		tx := c.Get("Tx").(*dbr.Tx)
		personInfo := new(model.PersonInfo)

		if err := personInfo.GetPersonInfo(tx, id); err != nil {
			//logrus.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound, "Resume does not exist.")
		}
		r.PersonInfo = *personInfo

		// education
		// ....


		return c.JSON(http.StatusOK, r)
	}
}*/