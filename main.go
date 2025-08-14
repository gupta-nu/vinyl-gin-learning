package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json: "title"`
	Artist string  `json: "artist"`
	Price  float64 `json: "price"`
}

//albums slice to seed record ablum data

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

//when client makes request get at /albums - we write
// logic to prepare a respone -> like someone asking wich cd's u have n u makign a list on a paper
//  code to map the request path to logic like if someone knocks on this door, which func will answer

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default() // initi gin rounter
	router.GET("/albums", getAlbums)
	router.Run("localhost:8080")
}
