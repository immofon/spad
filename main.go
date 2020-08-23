package main

import (
	"log"
	"net/http"
	"os"
)

func getenv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic("You must set env " + key)
	}
	return v
}

type FileSystem struct {
	IndexFile string
}

func (fs *FileSystem) Open(name string) (http.File, error) {
	f, err := os.Open(name)
	if err != nil {
		return os.Open(fs.IndexFile)
	}
	return f, nil
}

func main() {
	port := getenv("port")
	indexfile := getenv("indexfile")

	http.Handle("/", http.FileServer(&FileSystem{IndexFile: indexfile}))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Print(err)
	}
}
