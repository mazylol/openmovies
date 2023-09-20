package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Everything is working fine",
		})
	})

	movies, movieList, err := LoadMovies()
	if err != nil {
		log.Fatal("Failed to load movies")
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

		movie := movies[shortname]

		if movie.Shortname == "" {
			c.JSON(404, gin.H{
				"error": "Movie does not exist",
			})
		} else {
			c.JSON(200, movie)
		}

	})

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Could not open server")
	}
}
