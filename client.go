package main

import (
	"github.com/pkg/errors"
	"github.com/ypapax/tcp_ddos_golang/hashcash2"
	"log"
	"net"
	"time"
)
const PORT_CLIENT = "9001"
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := func() error {
		h := hashcash2.NewStd() // or .New(bits, saltLength, extra)
		// Mint a new stamp
		t1 := time.Now()
		stamp, err := h.Mint("client_id")
		if err != nil {
			return errors.WithStack(err)
		}
		log.Printf("time spent on generating the stamp: %+v", time.Since(t1))
		strEcho := stamp
		servAddr := "localhost:"+PORT_CLIENT
		tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
		if err != nil {
			return errors.WithStack(err)
		}

		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			return errors.WithStack(err)
		}

		_, err = conn.Write([]byte(strEcho))
		if err != nil {
			return errors.WithStack(err)
		}

		log.Println("write to server = ", strEcho)

		reply := make([]byte, 1024)

		_, err = conn.Read(reply)
		if err != nil {
			return errors.WithStack(err)
		}

		log.Println("reply from server=", string(reply))

		conn.Close()
		return nil
	}(); err != nil {
		log.Printf("error: %+v", err)
	}

}