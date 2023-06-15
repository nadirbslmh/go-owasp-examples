package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// User represents user entity in database
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

// UserInput represents request body for registration
type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// DB represents database simulation
var DB []User = []User{}

func main() {
	e := echo.New()

	// route for user registration
	e.POST("/auth/register", func(c echo.Context) error {
		var userInput UserInput

		if err := c.Bind(&userInput); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "register failed")
		}

		// create a password
		password, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "register failed")
		}

		// create user ID with UUID
		userID := uuid.NewString()

		user := User{
			ID:       userID,
			Email:    userInput.Email,
			Password: string(password),
		}

		// save to database
		DB = append(DB, user)

		return c.JSON(http.StatusCreated, echo.Map{
			"message": "register success",
		})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
