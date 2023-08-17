package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"nhooyr.io/websocket"
)

const (
	Host     = "localhost"
	Port     = "9988"
	Protocol = "tcp"
)

func main() {
	fmt.Println("Http Socker server starting on ...")
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			OriginPatterns: []string{"localhost"},
		})
		if err != nil {
			log.Println("Failed to accept connection: ", err.Error())
		}

		c.Close(websocket.StatusNormalClosure, "Your may leave now")
	})
	if err := http.ListenAndServe(":8080", fn); err != nil {
		fmt.Println("Failed with", err.Error())
		os.Exit(1)
	}
}
