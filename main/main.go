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
	MAX_SPEED float64 = 2
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

		liveCells := randPopulate(t.Width, t.Height, t.Density)
		board := trimCellTable(liveCells, t.Width, t.Height)

		err := c.WriteJSON(board)
		if err != nil {
			log.Println("ws send failed -- ", err)
			return
		}

		liveCells = turn(liveCells)
		board = trimCellTable(liveCells, t.Width, t.Height)

		speed := math.Max(t.Speed, MAX_SPEED)
		ticker := time.NewTicker(time.Millisecond * time.Duration(speed))

		defer ticker.Stop()

		for {
			<-ticker.C
			err := c.WriteJSON(board)
			liveCells = turn(liveCells)
			board = trimCellTable(liveCells, t.Width, t.Height)

			if err != nil {
				log.Println("write:", err)
				return
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", serveGame)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
