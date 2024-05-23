package cotacao

import (
	"fmt"
	"log"
	"os"
)

func WriteFile(value string) {
	file, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}

	length, err := file.WriteString(fmt.Sprintf("Dólar: %v", value))
	if err != nil {
		panic(err)
	}

	log.Printf("file length: %v", length)
}
