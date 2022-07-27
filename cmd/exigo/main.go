package main

import (
	"fmt"
	"net/http"

	"github.com/alwindoss/exigo"
	"github.com/caarlos0/env/v6"
)

func main() {
	fmt.Println("Welcome to Exigo")
	cfg := exigo.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})
	addr := fmt.Sprintf(":%d", cfg.Port)
	http.ListenAndServe(addr, nil)
}
