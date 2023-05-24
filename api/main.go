package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// Gerar token aleatorio
/* func init() {
	chave := make([]byte, 64)

	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	fmt.Println(stringBase64)
} */

func main() {
	config.Carregar()

	fmt.Printf("Escutando na porta %d", config.Porta)

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
