package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

// HeroData
type HeroData struct {
	HeroID     int // Hero yang dipilih
  HeroHP int // Health Point
	XPos int // Posisi X
	YPos int // posisi Y
  Side int // radiant/dire
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("Error : ", err)
	}
	fmt.Println("Server starting ...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	dec := gob.NewDecoder(conn)
	h := &HeroData{}
	dec.Decode(h)
	fmt.Println("Hero ID : ", h.HeroID," Hero HP: ", h.HeroHP ," Moving at x: ", h.XPos, " y: ", h.YPos, " Side : ", h.Side)
	conn.Close()
}
