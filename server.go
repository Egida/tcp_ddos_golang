package main

import (
	"github.com/pkg/errors"
	"log"
	"net"
	"os"
	"time"

	hc "github.com/catalinc/hashcash"
)

const (
	HOST = "localhost"
	PORT = "9001"
	TYPE = "tcp"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	log.Printf("listening %+v %+v:%+v ...", TYPE, HOST, PORT)
	// close listener
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go handleIncomingRequest(conn)
	}
}
func handleIncomingRequest(conn net.Conn) {
	// store incoming data
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	// respond
	time := time.Now().Format("Monday, 02-Jan-06 15:04:05 MST")
	conn.Write([]byte("Hi back!\n"))
	conn.Write([]byte(time))

	// close conn
	conn.Close()
}

func main2() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := func() error {
		h := hc.NewStd() // or .New(bits, saltLength, extra)

		// Mint a new stamp
		stamp1, err := h.Mint("something")
		if err != nil {
			return errors.WithStack(err)
		}
		stamp2, err := h.Mint("something")
		if err != nil {
			return errors.WithStack(err)
		}
		stamp7, err := h.Mint("something7")
		if err != nil {
			return errors.WithStack(err)
		}
		log.Printf("stamp1: %+v", stamp1)
		log.Printf("stamp2: %+v", stamp2)
		log.Printf("stamp7: %+v", stamp7)
		// Check a stamp
		valid := h.Check("1:20:161203:something::+YO19qNZKRs=:a31a2")
		if valid {
			log.Println("Valid")
		} else {
			log.Println("Invalid")
		}
		validNoDate := h.CheckNoDate("1:20:161203:something::+YO19qNZKRs=:a31a2")
		if validNoDate {
			log.Println("validNoDate: Valid")
		} else {
			log.Println("validNoDate: Invalid")
		}
		return nil
	}(); err != nil {
		log.Printf("err: %+v", err)
	}
	log.Println("finished")
}