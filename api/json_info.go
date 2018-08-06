package api

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/austinthale/resume-creator/model"
	"github.com/gocraft/dbr"
	"fmt"
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
	var r = model.Resume{}
	return func(c echo.Context) (err error) {
		tx := c.Get("Tx").(*dbr.Tx)
		personInfo := new(model.PersonInfo)
		if err = personInfo.Load(tx, 1); err != nil {
			fmt.Println(err)
		}
		r.PersonInfo = *personInfo
		return c.JSON(http.StatusOK, r)
	}
}