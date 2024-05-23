package main

import (
	"fullcycle/server/cotacao"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", cotacao.Handler)
	http.ListenAndServe(":8080", mux)
}
