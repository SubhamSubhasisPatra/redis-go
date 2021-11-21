package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type album struct {
	ID     string  `json:"id"`
	Artist string  `json:"artist"`
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
}

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var ctx = context.Background()

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var authers = []Author{
	{Name: "Gerry Mulligan", Age: 16},
	{Name: "Sarah Vaughan", Age: 23},
}

func postAuther(c *gin.Context) {
	var auther Author

	if err := c.BindJSON(&auther); err != nil {
		return
	}

	// pasrse the json
	json, err := json.Marshal(auther)
	if err != nil {
		panic(err)
	}
	err = client.HSet(ctx, "1", json, 0).Err()
	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, auther)
}

func getAutherById(c *gin.Context) {
	id := c.Param("id")
	// json, err := client.HGetAll(ctx, id).Result()
	json, err := client.Exists(ctx, id).Result()
	if err != nil {
		panic(err)
	}
	if json == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "auther not found"})
	} else {
		c.IndentedJSON(http.StatusOK, json)
	}
}

func getAlbum(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, newAlbum)
}

func getAlbumById(c *gin.Context) {

	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// Add the DB
var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func main() {
	router := gin.Default()
	// API routes for albums
	router.GET("/albums", getAlbum)
	router.POST("/albums", postAlbum)
	router.GET("/albums/:id", getAlbumById)

	// API routes for auther
	router.POST("/auther", postAuther)
	router.GET("/auther/:id", getAutherById)

	router.Run("localhost:8080")

}

/*

package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gitlab.com/idoko/rediboard/db"
)

var (
	ListenAddress = "locahost:8080"
	RedisAddress  = "localhost:6379"
)

func main() {
	database, err := db.NewDatebase(RedisAddress)
	if err != nil {
		log.Fatal(err.Error())
	}
	route := initRouter(database)
	route.Run(ListenAddress)
}

func initRouter(database *db.Database) *gin.Engine {
	r := gin.Default()
	return r
}

r.POST("/points",func (c *git.Context){

})

*/
