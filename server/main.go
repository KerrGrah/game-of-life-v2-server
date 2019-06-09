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

func read(c *websocket.Conn, tChan chan (GameSetup)) {
	t := GameSetup{}
	for {
		err := c.ReadJSON(&t)
		if err != nil {
			log.Print("read client failed -- ", err)
			return
		}
		fmt.Printf("Got message: %#v\n", t)
		tChan <- t
	}
}

func makeTicker(inputSpeed float64) *time.Ticker {
	speed := math.Max(inputSpeed, MAX_SPEED)
	return time.NewTicker(time.Millisecond * time.Duration(speed))
}

func serveGame(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("ws upgrade failed -- ", err)
		return
	}

	defer func() {
		err := c.Close()

		if err != nil {
			log.Print("ws upgrade failed -- ", err)
			return
		}
	}()

	liveCells := CellTable{}
	board := CellTable{}

	ticker := time.Ticker{}
	t := GameSetup{}
	tChan := make(chan GameSetup)

	go read(c, tChan)

	for {
		select {
		case <-ticker.C:
			err := c.WriteJSON(board)
			liveCells = turn(liveCells)

			board = trimCellTable(liveCells, t.Width, t.Height)

			if err != nil {
				log.Println("write:", err)
				return
			}
		case t = <-tChan:
			if t.Initiate {
				liveCells = randPopulate(t.Width, t.Height, t.Density)
				board = trimCellTable(liveCells, t.Width, t.Height)

				err := c.WriteJSON(board)
				if err != nil {
					log.Println("error in write initial board state:", err)
					return
				}
				ticker = *makeTicker(t.Speed)
				defer ticker.Stop()

				liveCells = turn(liveCells)
			} else {
				ticker = *makeTicker(t.Speed)
			}
		}
	}

}
func main() {
	buildHandler := http.FileServer(http.Dir("../client/build"))

	http.Handle("/", buildHandler)
	http.HandleFunc("/ws", serveGame)

	log.Fatal(http.ListenAndServe(*addr, nil))
}
