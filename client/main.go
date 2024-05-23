package main

import (
	"client/cotacao"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type CotacaoResponse struct {
	Dolar string `json:"Bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var cotacaoResponse CotacaoResponse
	err = json.Unmarshal(body, &cotacaoResponse)
	if err != nil {
		panic(err)
	}

	cotacao.WriteFile(cotacaoResponse.Dolar)
}
