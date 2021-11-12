package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	json, err := json.Marshal(Author{Name: "Subhasis", Age: 19})
	fmt.Println(string(json))
	if err != nil {
		panic(err)
	}
	// set the value
	err = client.HSet(ctx, "person", json, 0).Err()
	// err = client.HSet(ctx, "address", map[string]interface{}{"phone": "7325902451", "state": "Odisha"}).Err()
	if err != nil {
		panic(err)
	}

	// Get the Value
	val, err := client.HGetAll(ctx, "person").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("person", val)

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
