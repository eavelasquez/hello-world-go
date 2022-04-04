package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "4", Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Price: 8.99},
}

// health check endpoint to check if the service is up and running
func healthCheck(c *gin.Context) {
	// Set a 200 (OK) response code.
	c.IndentedJSON(http.StatusOK, gin.H{"status": "ok"})
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	// Return the list of albums as JSON.
	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	// Get the ID parameter from the request.
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			// If found, return the album as JSON.
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	// If no album was found, return a status 404.
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to the
	// newAlbum variable.
	if err := c.BindJSON(&newAlbum); err != nil {
		// If there is an error, return a 400 (Bad Request)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add the new album to the slice of albums.
	albums = append(albums, newAlbum)

	// Return the new album as JSON.
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// updateAlbum updates an album with new data.
func updateAlbum(c *gin.Context) {
	// Get the ID parameter from the request.
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID matches the parameter.
	for index, a := range albums {
		if a.ID == id {
			// If found, update the album.
			var updatedAlbum album

			// Call BindJSON to bind the received JSON to the
			// updatedAlbum variable.
			if err := c.BindJSON(&updatedAlbum); err != nil {
				// If there is an error, return a 400 (Bad Request)
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// Set the ID to the same ID sent in the request.
			updatedAlbum.ID = a.ID

			// Update the album and replace it in the slice.
			albums[index] = updatedAlbum

			// Return the updated album as JSON.
			c.IndentedJSON(http.StatusOK, albums[index])
			return
		}
	}
}

// deleteAlbum removes an album from the list.
func deleteAlbum(c *gin.Context) {
	// Get the ID parameter from the request.
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID matches the parameter.
	for index, a := range albums {
		if a.ID == id {
			// If found, remove the album.
			albums = append(albums[:index], albums[index+1:]...)

			// Return the deleted album as JSON.
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	// If no album was found, return a status 404.
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Album not found"})
}

func main() {
	// Create an instance of the router.
	router := gin.Default()

	// Set a route group for the API.
	router.GET("/", healthCheck)              // Health check.
	router.GET("/albums", getAlbums)          // List all albums.
	router.GET("/albums/:id", getAlbumByID)   // Get an album by it's ID.
	router.POST("/albums", postAlbums)        // Create a new album.
	router.PUT("/albums/:id", updateAlbum)    // Update an album.
	router.DELETE("/albums/:id", deleteAlbum) // Delete an album.

	// Listen and Server in
	router.Run("localhost:8080")
}
