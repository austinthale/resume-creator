package api

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/austinthale/resume-creator/model"
)
// TODO build the r variable by accessing database and convert to structs

func DisplayInfo() echo.HandlerFunc {
	var r = model.Resume{}
	return func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, r)
	}
}

func SaveInfo() echo.HandlerFunc {
	var r = model.Resume{}
	return func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, r)
	}
}

func DisplayJSON() echo.HandlerFunc {
	//var r = model.Resume{}
	return func(c echo.Context) (err error) {
		//tx := c.Get("Tx").(*dbr.Tx)
		personInfo := new(model.PersonInfo)
		// TODO ERROR OCCURS HERE
		/*if err := personInfo.Load(tx, 1); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound, "Resume does not exist.")
		}*/
		//r.PersonInfo = *personInfo
		return c.JSON(http.StatusOK, personInfo)
	}
}