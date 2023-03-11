package main

import (
	"bytes"
	"github.com/ypapax/tcp_ddos_golang/hashcash2"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
)

const (
	HOST = "" // localhost
	PORT = "9001"
	TYPE = "tcp"
	jokesFile = "jokes.txt"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	h := hashcash2.NewStd()
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
		var (
			uniqueMap = make(map[string]time.Time)
			uniqueMapMtx = sync.Mutex{}
		)
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
			fromClient := strings.TrimSpace(string(bytes.Trim(buffer, "\x00")))
			log.Printf("received from client: %+v, len(fromClient): %+v",
				fromClient, len(fromClient))
			writeToClient := func(msg string) error {
				if _, errW := conn.Write([]byte(msg)); errW != nil {
					return errors.WithStack(errW)
				}
				return nil
			}
			var msgToClient string
			t, found := uniqueMap[fromClient] // better to use database to keep date between server reruns
			if found {
				log.Printf("this token was already detected in the past %+v : %+v", fromClient, t)
			}
			// TODO: add tests and check how this method works
			if found || !h.Check(fromClient) {
				msgToClient = "the request is not verified by proof of work hashcash"
			} else {
				l, errR := randomLine()
				if errR != nil {
					return errors.WithStack(errR)
				}
				msgToClient = l
			}
			func(){
				uniqueMapMtx.Lock()
				defer uniqueMapMtx.Unlock()
				uniqueMap[fromClient] = time.Now()
			}()
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
		log.Printf("listening %+v %+v:%+v ....", TYPE, HOST, PORT)
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