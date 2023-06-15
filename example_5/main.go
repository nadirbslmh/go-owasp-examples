package main

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(c echo.Context) error {
	var input UserInput

	if err := c.Bind(&input); err != nil {
		return echo.ErrBadRequest
	}

	if input.Username != "admin" || input.Password != "admin" {
		return echo.ErrUnauthorized
	}

	claims := &jwtCustomClaims{
		"admin",
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func getUsers(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

// verifyRole checks user role
func verifyRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)
		isAdmin := claims.Admin

		if !isAdmin {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}

func main() {
	e := echo.New()

	// Login route
	e.POST("/login", login)

	// Restricted group for admin only
	adminRoutes := e.Group("/users")

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}

	adminRoutes.Use(echojwt.WithConfig(config))

	// check if logged in user is admin
	adminRoutes.Use(verifyRole)

	adminRoutes.GET("", getUsers)

	e.Logger.Fatal(e.Start(":1323"))
}
