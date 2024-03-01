package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/josebruno2020/go-devbook-api/src/config"
	"github.com/josebruno2020/go-devbook-api/src/db"
	"github.com/josebruno2020/go-devbook-api/src/router"
)

func main() {
	config.Setup()
	db.Connect()
	fmt.Println(config.ConnectionDB)
	port := config.AppPort
	fmt.Printf("Rodando na porta %d\n", port)
	r := router.Setup()

	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%d", port),
			r,
		),
	)
}
