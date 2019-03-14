package main

import (
    "encoding/gob"
  	"fmt"
  	"log"
  	"net"
    "os"
)

type HeroData struct {
	HeroID     int // Hero yang dipilih
  HeroHP int // Health Point
	XPos int // Posisi X
	YPos int // posisi Y
  Side int // radiant/dire
}

func main() {
	fmt.Println("NetClient Encap")
	//create structure object

	fmt.Println("Client Starting ...")
	// dial TCP connection
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("Connection error", err)
    os.exit(1);
	}
	//Create encoder object, We are passing connection object in Encoder
	// Encode Structure, IT will pass object over TCP connection
  for {
		heroData := HeroData{ID: 22,HeroHP: 78, XPos: 31, YPos: 224, Side: 2}
		encoder := gob.NewEncoder(conn)
		encoder.Encode(heroData)

		conn.Close()
}


	// close connection
	conn.Close()
	fmt.Println("Exit")
}
