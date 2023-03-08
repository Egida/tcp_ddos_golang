package main

import (
	"log"
	"net"

	hc "github.com/catalinc/hashcash"
	"github.com/pkg/errors"
)

const (
	HOST = "localhost"
	PORT = "9001"
	TYPE = "tcp"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	h := hc.NewStd()
	handleIncomingRequest := func(conn net.Conn) error {
		// store incoming data
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			return errors.WithStack(err)
		}
		fromClient := string(buffer)
		log.Printf("received from client: %+v", fromClient)
		writeToClient := func(msg string) error {
			if _, errW := conn.Write([]byte(msg)); errW != nil {
				return errors.WithStack(errW)
			}
			return nil
		}
		var msgToClient string
		if h.Check(fromClient) {
			msgToClient = "the request is not verified by proof of work hashcash"
		} else {
			msgToClient = "ok"
		}
		if errW := writeToClient(msgToClient); errW != nil {
			return errors.WithStack(errW)
		}


		// close conn
		if errC := conn.Close(); errC != nil {
			return errors.WithStack(errC)
		}
		return nil
	}
	if err := func() error {
		listen, err := net.Listen(TYPE, HOST+":"+PORT)
		if err != nil {
			return errors.WithStack(err)
		}
		log.Printf("listening %+v %+v:%+v ...", TYPE, HOST, PORT)
		// close listener
		defer listen.Close()
		for {
			conn, errL := listen.Accept()
			if errL != nil {
				return errors.WithStack(errL)

			}
			go func() {
				if errH := handleIncomingRequest(conn); errH != nil {
					log.Printf("couldn't process client request: %+v", errH)
				}
			}()
		}
	}(); err != nil {
		log.Printf("error: %+v", err)
	}

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
