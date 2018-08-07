package api

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/austinthale/resume-creator/model"
	"github.com/gocraft/dbr"
	"strconv"
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

		// Load Personal Info
		personInfo := new(model.PersonInfo)
		if err := personInfo.Load(tx, id); err != nil {
			//logrus.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound, "Resume does not exist.")
		}
		r.PersonInfo = *personInfo

		// Load data into resume
		r.Load(tx, id)

		return c.JSON(http.StatusOK, r)
	}
}

func SaveInfo() echo.HandlerFunc {
	var r = model.Resume{}
	return func(c echo.Context) (err error) {
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

		if err := personInfo.Load(tx, id); err != nil {
			//logrus.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound, "Resume does not exist.")
		}
		r.PersonInfo = *personInfo

		// education
		// ....


		return c.JSON(http.StatusOK, r)
	}
}*/