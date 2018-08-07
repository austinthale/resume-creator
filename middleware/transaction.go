package middleware

import (
	"github.com/sirupsen/logrus"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
)

const (
	TxKey = "Tx"
)

func TransactionHandler(db *dbr.Session) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			logrus.Println(db)
			tx, err := db.Begin()
			logrus.Println("Tx value: ", tx)
			c.Set(TxKey, tx)
			logrus.Println("TxKey: ", TxKey, "tx: ", tx, err)

			// If there's an error, Rollback
			if err := next(c); err != nil {
				tx.Rollback()
				logrus.Debug("Transction Rollback: ", err)
				return err
			}
			// else, Commit transaction
			logrus.Debug("Transaction Commit")
			tx.Commit()

			return nil
		})
	}
}
