package server

import (
	"errors"
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var server *http.Server

type FileSystem struct {
	fs http.FileSystem
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}

func Start() {

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	directory := flag.String("d", dir+"/views", "The directory of static file to host")

	fs := http.FileServer(http.FileSystem(http.Dir(*directory)))

	http.Handle("/", http.StripPrefix(strings.TrimRight("/", "/"), fs))

	server = &http.Server{
		Addr:           ":5432",
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(errors.New(fmt.Sprintf("Server shutdown with fatal error: %s", err)))
		}
	}()
}

func Stop() {
	//Shutdown server
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown with error: %s", err.Error())
	}
}
