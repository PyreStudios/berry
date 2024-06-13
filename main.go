package main

import (
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("sample_config.yml")
	viper.ReadInConfig()

	indexablePaths := viper.GetStringSlice("media_roots")

	var files []Media
	for _, path := range indexablePaths {
		files = append(files, LoadMediaFilesFromPath(path)...)
	}

	http.HandleFunc("GET /media/comics", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Get all comics
	})

	http.HandleFunc("GET /media/books", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Get all books
	})

}
