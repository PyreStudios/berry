package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type MediaType int

const (
	MediaTypeUnknown = iota
	MediaTypeBook
	MediaTypeComic
)

type Media struct {
	Path           string
	Extension      string
	MediaType      MediaType
	CoverImagePath string
}

func LoadMediaFilesFromPath(rootPath string) []Media {
	var files []Media
	err := filepath.Walk(rootPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		pathWithExt := strings.Split(path, ".")

		if len(path) == 1 {
			return errors.New("file did not contain a valid extension")
		}

		mediaType := MediaTypeUnknown

		if strings.HasPrefix(path, filepath.FromSlash(rootPath+"/books")) {
			mediaType = MediaTypeBook
		} else if strings.HasPrefix(path, filepath.FromSlash(rootPath+"/comics")) {
			mediaType = MediaTypeComic
		} else if strings.HasPrefix(path, filepath.FromSlash(rootPath+"/manga")) {
			mediaType = MediaTypeComic
		}

		var coverImagePath string = pathWithExt[0] + "_cover.jpg"
		if _, err := os.Stat(coverImagePath); err == nil {
			coverImagePath = ""
		}

		files = append(files, Media{
			Path:           path,
			Extension:      pathWithExt[len(pathWithExt)-1],
			MediaType:      MediaType(mediaType),
			CoverImagePath: coverImagePath,
		})
		return nil
	})

	if err != nil {
		// TODO: Lets add more here
		fmt.Println("Error parsing certain files.")
	}

	return files
}
