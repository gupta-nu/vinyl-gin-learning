package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
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
	//statusok is req succeded and i'm returning the data used for GET req
}

//logic to add new album into existing list
//code to route the post request to logic

func postAlbums(c *gin.Context) {
	var newAlbum album //bindJSON rads the json payload from http req
	//and map(bind) it to go variable(newalubm)
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	} //err checking, if err!=nil, there was error and jscon couldn't be parsed;

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
	//statusCreated is for pos requests that add new resources
}

//logic to return specific item ,to retrieve requested album ,
//to map the path to logic

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a) //sends back a 200 ok status to client, and returns a
			return
		}
	}
	//if no match, 404 respone it send and
	//gin.H is map[string]interface{} , - its a go map, key-string, value- anything,
	//so we send gin.H of the message of album not found
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default() // initi gin rounter
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
