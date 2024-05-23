package cotacao

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type UsdBrl struct {
	Code       string
	Codein     string
	Name       string
	High       string
	Low        string
	VarBid     string
	PctChange  string
	Bid        string `json:"Bid"`
	Ask        string
	Timestamp  string
	CreateDate string `json:"create_date"`
}

type UsdBrlResponse struct {
	Bid string
}

type CotacaoDolar struct {
	UsdBrl UsdBrl
}

func Handler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var cotacao CotacaoDolar
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		panic(err)
	}

	response := UsdBrlResponse{
		Bid: cotacao.UsdBrl.Bid,
	}
	//salvar no banco de dados sqlite.
	json.NewEncoder(w).Encode(response)
}
