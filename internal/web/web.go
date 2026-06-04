package web

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

//go:embed dist/**
var webFiles embed.FS

func StartWebServer() error {
	subFS, err := fs.Sub(webFiles, "dist")
	if err != nil {
		return err
	}

	fileServer := http.FileServer(http.FS(subFS))
	http.Handle("/", spaHandler(fileServer))
	fmt.Println("Web Server is served at localhost:9056")

	return http.ListenAndServe(":9056", nil)
}

func spaHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fs.Stat(webFiles, "dist"+r.URL.Path)
		if err == nil {
			h.ServeHTTP(w, r)
			return
		}

		data, err := webFiles.ReadFile("dist/index.html")
		if err != nil {
			http.Error(w, "Not Found", 404)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		w.Write(data)
	})
}
