package api

import (
	"github.com/labstack/echo"
	"github.com/austinthale/resume-creator/model"
	"github.com/gocraft/dbr"
	"github.com/sirupsen/logrus"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func Login() echo.HandlerFunc{
	return func(c echo.Context) (err error) {
		user := new(model.User)

		tx := c.Get("Tx").(*dbr.Tx)

		// Bind gets the JSON data from the page and binds it to the Resume object
		if err = c.Bind(user); err != nil {
			logrus.Println("Error binding data: ", err)
			return
		}
		u := *user 	// 'u' now reflects the input that we used BIND to get, we can start validating
		// Validate username and password
		if err := u.Validate(tx); err != nil {
			logrus.Println("Error Validating: ", err)
			return c.NoContent(http.StatusUnauthorized)
		} else {
			// Create token
			token := jwt.New(jwt.SigningMethodHS256)

			// Set claims
			claims := token.Claims.(jwt.MapClaims)
			claims["name"] = u.Username
			claims["admin"] = true
			claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

			// Generate encoded token and send it as response.
			t, err := token.SignedString([]byte("secret"))
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, map[string]string{
				"token": t,
			})
		}
	}
}

func Accessible() echo.HandlerFunc {
	return func(c echo.Context) error{
		return c.String(http.StatusOK, "Accessible")
	}
}

func Restricted() echo.HandlerFunc {
	return func (c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
	}
}


