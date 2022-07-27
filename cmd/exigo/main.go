package main

import (
	"fmt"
	"net/http"

	"github.com/alwindoss/exigo"
	"github.com/caarlos0/env/v6"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to Exigo")
	cfg := exigo.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg)
	r := mux.NewRouter()

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	// originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// spa := spaHandler{staticFS: ui.DistDir, staticPath: "dist", indexPath: "index.html"}
	// r.PathPrefix("/").Methods(http.MethodGet).Handler(spa)
	r.Path("/api/ping").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})
	r.Path("/api/search").Methods(http.MethodGet).HandlerFunc(searchHandler)
	addr := fmt.Sprintf(":%d", cfg.Port)
	http.ListenAndServe(addr, handlers.CORS(originsOk, headersOk, methodsOk)(r))
}
