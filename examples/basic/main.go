package main

import (
	"fmt"
	"net/http"

	"github.com/domudall/go-binance"
)

func main() {
	httpClient := http.DefaultClient
	client := binance.NewClient(httpClient)

	err := client.Ping()
	fmt.Println("ping:", err)

	serverTime, err := client.Time()
	fmt.Println("time:", serverTime, err)
}
