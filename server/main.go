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

func read(c *websocket.Conn, t GameSetup) GameSetup {
	err := c.ReadJSON(&t)
	if err != nil {
		fmt.Println("Error reading json.", err)
	}

	fmt.Printf("Got message: %#v\n", t)
	return t
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

	defer c.Close()

	t := GameSetup{}

	t = read(c, t)

	tChan := make(chan GameSetup)

	go func() {
		for {
			tChan <- read(c, t)
		}
	}()

	liveCells := randPopulate(t.Width, t.Height, t.Density)
	board := trimCellTable(liveCells, t.Width, t.Height)

	ticker := makeTicker(t.Speed)
	defer ticker.Stop()

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
			ticker = makeTicker(t.Speed)
		}
	}

}
func main() {
	buildHandler := http.FileServer(http.Dir("../client/build"))

	http.Handle("/", buildHandler)
	http.HandleFunc("/ws", serveGame)

	log.Fatal(http.ListenAndServe(*addr, nil))
}
