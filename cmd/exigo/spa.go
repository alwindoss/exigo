package main

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
)

// //go:embed build/*
// var staticFiles embed.FS

type spaHandler struct {
	staticFS   embed.FS
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	_, err = h.staticFS.Open(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		index, err := h.staticFS.ReadFile(filepath.Join(h.staticPath, h.indexPath))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusAccepted)
		w.Write(index)
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get the subdirectory of the static dir
	statics, err := fs.Sub(h.staticFS, h.staticPath)
	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.FS(statics)).ServeHTTP(w, r)
}

// func main() {
// 	r := mux.NewRouter()
// 	spa := spaHandler{staticFS: staticFiles, staticPath: "build", indexPath: "index.html"}
// 	r.PathPrefix("/").Handler(spa)
// ...
// }

// func main() {
// 	router := mux.NewRouter()

// 	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
// 		// an example API handler
// 		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
// 	})

// 	spa := spaHandler{staticPath: "build", indexPath: "index.html"}
// 	router.PathPrefix("/").Handler(spa)

// 	srv := &http.Server{
// 		Handler: router,
// 		Addr:    "127.0.0.1:8000",
// 		// Good practice: enforce timeouts for servers you create!
// 		WriteTimeout: 15 * time.Second,
// 		ReadTimeout:  15 * time.Second,
// 	}

// 	log.Fatal(srv.ListenAndServe())
// }
