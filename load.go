package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mazylol/openmovies/types"
)

func LoadMovies() (map[string]types.Movie, []string, error) {
	var movies = make(map[string]types.Movie)

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

		if err != nil {
			log.Fatal("Failed to unmashal json data for")
		}

		filename, _, _ := strings.Cut(file.Name(), ".")

		movies[filename] = movie
	}

	var movieList = make([]string, 0, len(movies))

	for key := range movies {
		movieList = append(movieList, key)
	}

	return movies, movieList, err
}
