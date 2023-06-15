package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Data struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    Book   `json:"data"`
}

func main() {
	e := echo.New()

	e.GET("/resources", func(c echo.Context) error {
		book := Book{
			ID:    1,
			Title: "sample",
		}

		data := Data{
			Status:  true,
			Message: "book data",
			Data:    book,
		}

		return c.JSON(http.StatusOK, data)
	})

	e.Logger.Fatal(e.Start(":5000"))
}
