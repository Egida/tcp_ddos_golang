package main

import (
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"

	hc "github.com/catalinc/hashcash"
	"github.com/pkg/errors"
)

const (
	HOST = "localhost"
	PORT = "9001"
	TYPE = "tcp"
	jokesFile = "jokes.txt"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	h := hc.NewStd()
	if err := func() error {
		b, err := os.ReadFile(jokesFile)
		if err != nil {
			return errors.WithStack(err)
		}
		lines := strings.Split(string(b), "\n")
		if len(lines) <= 1 {
			return errors.Errorf("to few jokes")
		}
		log.Printf("loaded %+v jokes", len(lines))
		rand.Seed(time.Now().Unix())
		randomLine := func() (string, error) {
			lineN := rand.Intn(len(lines))
			if lineN >= len(lines) {
				return "", errors.Errorf("not valid line number")
			}
			line := lines[lineN]
			if len(line) == 0 {
				return line, errors.Errorf("to short line")
			}
			return line, nil
		}
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
				l, errR := randomLine()
				if errR != nil {
					return errors.WithStack(errR)
				}
				msgToClient = l
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