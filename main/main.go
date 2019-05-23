package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math"
	"net/http"
	"time"
)

var (
	MAX_SPEED float64 = 10
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func serveGame(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool {
		// TODO add check
		return true
	}

	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("ws upgrade failed -- ", err)
		return
	}

	defer c.Close()

	for {
		t := GameSetup{}
		err = c.ReadJSON(&t)

		if err != nil {
			fmt.Println("Error reading json.", err)
		}

		fmt.Printf("Got message: %#v\n", t)

		board := randPopulate(t.Height, t.Width, t.Density)

		for {
			err := c.WriteJSON(board)
			if err != nil {
				log.Println("ws send failed -- ", err)
				return
			}
			board = turn(board)

			speed := math.Max(t.Speed, MAX_SPEED)

			time.Sleep(time.Duration(speed) * time.Millisecond)
		}
	}
}

func main() {
	http.HandleFunc("/ws", serveGame)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
