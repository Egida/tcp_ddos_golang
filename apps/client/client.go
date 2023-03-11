package main

import (
	"github.com/pkg/errors"
	"github.com/ypapax/tcp_ddos_golang/hashcash2"
	"log"
	"net"
	"os"
	"time"
)
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := func() error {
		servAddr := os.Getenv("SERVER_ADDR")
		if len(servAddr) == 0 {
			return errors.Errorf("missing server addr")
		}
		h := hashcash2.NewStd() // or .New(bits, saltLength, extra)
		
		tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
		if err != nil {
			return errors.WithStack(err)
		}
		for i:=0;i<=10;i++ {
			if errF := func() error {
				conn, errD := net.DialTCP("tcp", nil, tcpAddr)
				if errD != nil {
					return errors.WithStack(errD)
				}
				defer conn.Close()
				t1 := time.Now()
				stamp, err := h.Mint("client_id")
				if err != nil {
					return errors.WithStack(err)
				}
				log.Printf("time spent on generating the stamp: %+v", time.Since(t1))
				strEcho := stamp
				if _, errW := conn.Write([]byte(strEcho)); errW != nil {
					return errors.WithStack(errW)
				}
				log.Println("write to server = ", strEcho)
				reply := make([]byte, 1024)
				_, errD = conn.Read(reply)
				if errD != nil {
					return errors.WithStack(errD)
				}
				log.Println("reply from server=", string(reply))
				sl := 10 * time.Second
				log.Printf("sleeping for %+v", sl)
				time.Sleep(sl)
				return nil
			}(); errF != nil {
				return errors.WithStack(errF)
			}
		}
		return nil
	}(); err != nil {
		log.Printf("error: %+v", err)
	}

}