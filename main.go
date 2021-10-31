package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nickrisaro/freezer-bot/telegram"
)

func main() {
	log.Print("Iniciando freezer-bot")

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	token := os.Getenv("TELEGRAM_API_TOKEN")
	urlPublica := os.Getenv("APP_URL")

	if host == "" {
		log.Print("Usando host default")
		host = "0.0.0.0"
	}

	if port == "" {
		log.Print("Usando port default")
		port = "3000"
	}

	if token == "" {
		log.Fatal("Se debe setear la variable TELEGRAM_API_TOKEN")
		return
	}

	if urlPublica == "" {
		log.Fatal("Se debe setear la variable APP_URL")
		return
	}

	b, err := telegram.Configurar(urlPublica, fmt.Sprintf("%s:%s", host, port), token)

	if err != nil {
		log.Fatal("No pude iniciar el bot", err)
		return
	}

	log.Printf("Webhook escuchando en %s -> %s:%s", urlPublica, host, port)
	b.Start()
}
