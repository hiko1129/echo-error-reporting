package reporting

import (
	"log"

	"cloud.google.com/go/errorreporting"
	"github.com/labstack/echo/v4"
)

// New returns a new echo middleware which sends an error to gcp error reporting
func New(client *errorreporting.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)

			if err == nil {
				return nil
			}

			log.Println(err)

			client.Report(errorreporting.Entry{
				Error: err,
			})

			return err
		}
	}
}
