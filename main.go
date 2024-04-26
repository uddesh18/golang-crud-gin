package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Person struct {
	Name    string `json:"name"`
	Address string `json:"add"`
	Age     int    `json:"age"`
	Gender  string `json:"gender"`
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {

	person1 := &Person{"Uddesh", "Nanded", 24, "M"}
	a, _ := json.Marshal(person1)

	var albums = []album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	log.Info().Msg("Server Starting")

	routes := gin.Default()

	routes.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, string(a))
	})

	routes.GET("/allAlbums", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, albums)
	})

	routes.POST("/albumId", func(ctx *gin.Context) {
		id := ctx.Param("id")

		for _, a := range albums {
			if a.ID == id {
				ctx.IndentedJSON(http.StatusOK, a)
				return
			}
		}
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	})

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()

	fmt.Println(err)

}
