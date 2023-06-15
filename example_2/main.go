package main

import (
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
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
	Email string `json:"email" validate:"required,email"`
	// minimum 8 characters with 1 special character dan 1 number character
	Password string `json:"password" validate:"required,min=8,containsNumber,containsSpecialCharacter"`
}

func (input *UserInput) Validate() error {
	validate := validator.New()

	validate.RegisterValidation("containsNumber", func(fl validator.FieldLevel) bool {
		match, _ := regexp.MatchString(`[0-9]`, fl.Field().String())
		return match
	})
	validate.RegisterValidation("containsSpecialCharacter", func(fl validator.FieldLevel) bool {
		match, _ := regexp.MatchString(`[!@#$%^&*()]`, fl.Field().String())
		return match
	})

	err := validate.Struct(input)

	return err
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

		// validate request
		err := userInput.Validate()

		if err != nil {
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
