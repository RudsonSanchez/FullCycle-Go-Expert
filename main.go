package main

import (
	"net/http"
	"server/cotacao"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", cotacao.Handler)
	http.ListenAndServe(":8080", mux)
}
