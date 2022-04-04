package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// user represents data about a record user.
type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*user{} // User in-memory database
	seq   = 1               // Sequence
)

// bindUsers converts a map to a slice of users.
func bindUsers(m map[int]*user) []user {
	// Create slice
	us := make([]user, len(m))
	i := 0
	// Loop through map
	for _, u := range m {
		us[i] = *u
		i++
	}

	return us
}

// helthcheck is a simple healthcheck endpoint.
func helthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

// getUsers returns a list of users.
func getUsers(c echo.Context) error {
	// Get users
	return c.JSON(http.StatusOK, bindUsers(users))
}

// getUserByID returns a user for a given ID.
func getUserByID(c echo.Context) error {
	// Get ID from path
	id, _ := strconv.Atoi(c.Param("id"))

	// Get user
	u, ok := users[id]
	if !ok {
		return echo.ErrNotFound
	}

	// Response
	return c.JSON(http.StatusOK, u)
}

// createUser creates a new user.
func createUser(c echo.Context) error {
	u := &user{
		ID: seq,
	}

	// Bind user data from request
	if err := c.Bind(u); err != nil {
		return err
	}

	// Validate name
	if u.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "name cannot be empty")
	}

	// Save user
	users[u.ID] = u
	seq++

	// Response
	return c.JSON(http.StatusCreated, u)
}

// updateUser updates an existing user.
func updateUser(c echo.Context) error {
	// Get ID from path
	id, _ := strconv.Atoi(c.Param("id"))

	// Get user
	u, ok := users[id]
	if !ok {
		return echo.ErrNotFound
	}

	// Bind user data from request
	if err := c.Bind(u); err != nil {
		return err
	}

	// Validate name
	if u.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "name cannot be empty")
	}

	// Save user
	users[u.ID] = u

	// Response
	return c.JSON(http.StatusOK, u)
}

// deleteUser deletes an existing user.
func deleteUser(c echo.Context) error {
	// Get ID from path
	id, _ := strconv.Atoi(c.Param("id"))

	// Get user
	_, ok := users[id]
	if !ok {
		return echo.ErrNotFound
	}

	// Delete user
	delete(users, id)

	// Response
	return c.NoContent(http.StatusNoContent)
}

// main function is the entry point for the application
func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", helthcheck)
	e.GET("/users", getUsers)
	e.GET("/users/:id", getUserByID)
	e.POST("/users", createUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
