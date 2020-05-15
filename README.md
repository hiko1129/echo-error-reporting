# echo-error-reporting
echo-error-reporting is echo middleware which sends an error to gcp error reporing

## Installation
```
go get -u https://github.com/hiko1129/echo-error-reporting
```

## Usage
```
package main

import (
	"context"
	"errors"
	"net/http"
	"os"

	"cloud.google.com/go/errorreporting"
	reporting "github.com/hiko1129/echo-error-reporting"
	"github.com/labstack/echo/v4"
)

func example() {
	ctx := context.Background()
	errorClient, err := errorreporting.NewClient(ctx, os.Getenv("PROJECT_ID"), errorreporting.Config{})

	if err != nil {
		panic(err)
	}

	defer errorClient.Close()

	e := echo.New()
	e.Use(reporting.New(errorClient)) // set middleware

	// return error
	e.GET("/example", func(c echo.Context) error {
		err := errors.New("foo")

		jsonErr := c.JSON(http.StatusBadRequest, map[string]interface{}{
			"errorMessage": err.Error(),
		})

		if jsonErr != nil {
			return jsonErr
		}

		return err
	})

	e.Logger.Fatal(e.Start(":8080"))
}
```

echo-error-reporting sends an error to gcp error reporting if `echo.HandlerFunc` returns an error
