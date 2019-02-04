package main

import (
	"github.com/BlckStar/user"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())

	// Routes
	e.GET("/", func(c echo.Context) error { return hello(c, e) })
	e.GET("/users", listUsers)
	e.GET("/users/:user", getUser)
	e.POST("/users", addUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context, e *echo.Echo) error {
	return c.JSON(http.StatusOK, e.Routes())
}

func listUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, user.GetAllUsers());
}

func getUser(c echo.Context) error {
	var user, err =  user.GetUser(c.Param("user"))
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func addUser(c echo.Context) error {
	u := new(user.User)
	var err = c.Bind(u)

	if err != nil {
    	return c.String(http.StatusBadRequest, err.Error())
	}

	var userErr = user.AddUser(u);

	if userErr != nil {
    	return c.String(http.StatusBadRequest, userErr.Error())
	}
	  
  	return c.JSON(http.StatusOK, u)
}
