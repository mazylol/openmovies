package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mazylol/openmovies/types"
	"log"
	"os"
	"strings"
)

var movies = make(map[string]types.Movie)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Everything is working fine",
		})
	})

	files, err := os.ReadDir("content/movies")
	if err != nil {
		log.Fatal("Could not read content/movies")
	}

	for _, file := range files {
		bytes, err := os.ReadFile(fmt.Sprintf("content/movies/%s", file.Name()))
		if err != nil {
			log.Fatal("Could not read file contents")
		}

		var movie types.Movie

		err = json.Unmarshal(bytes, &movie)

		filename, _, _ := strings.Cut(file.Name(), ".")

		movies[filename] = movie
	}

	var movieList = make([]string, 0, len(movies))

	for key := range movies {
		movieList = append(movieList, key)
	}

	rg := r.Group("/api")

	rg.GET("/movie", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"count":  len(movieList),
			"movies": movieList,
		})
	})

	rg.GET("/movie/:shortname", func(c *gin.Context) {
		shortname := c.Param("shortname")

		c.JSON(200, movies[shortname])
	})

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Could not open server")
	}
}
