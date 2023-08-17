package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"nhooyr.io/websocket"
)

const (
	Host              = "localhost"
	Port              = "9988"
	Protocol          = "tcp"
	FIRST_MSG_TIMEOUT = time.Second * 20
)

func main() {
	fmt.Println("Http Socker server starting on ...")
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			OriginPatterns: []string{"localhost"},
			Subprotocols:   []string{"echo"},
		})
		if err != nil {
			log.Println("Failed to accept connection: ", err.Error())
		}

		fmt.Println("Going to handle")
		go handleSocketMessage(c, w, r)
		//c.Close(websocket.StatusNormalClosure, "Your may leave now")
	})
	if err := http.ListenAndServe(":8080", fn); err != nil {
		fmt.Println("Failed with", err.Error())
		os.Exit(1)
	}
}

func handleSocketMessage(c *websocket.Conn, w http.ResponseWriter, r *http.Request) {
	fmt.Println("About to handle")
	readTimeoutCtx, cancelRead := context.WithTimeout(context.Background(), FIRST_MSG_TIMEOUT)
	defer cancelRead()

	msgType, data, err := c.Read(readTimeoutCtx)
	if err != nil {
		c.Close(websocket.StatusNormalClosure, "Are you shy?")
		fmt.Println("Failed to read: ", err.Error())
		return
	}
	fmt.Printf("SP %v\n", c.Subprotocol())
	fmt.Printf("MT %v\n", msgType)
	fmt.Printf("Data %v\n", string(data))

	c.Close(websocket.StatusNormalClosure, "Your may leave now")
}
