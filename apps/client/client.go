package main

import (
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/ypapax/tcp_ddos_golang/common"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := func() error {
		servAddr := os.Getenv("SERVER_ADDR")
		if len(servAddr) == 0 {
			return errors.Errorf("missing server addr")
		}
		h, err := common.HashcashObjFromEnv()
		if err != nil {
			return errors.WithStack(err)
		}
		const clientReqsAmount = 10
		for i := 0; i < clientReqsAmount; i++ {
			if errF := func() error {
				t1 := time.Now()
				stamp, errM := h.Mint("client_id") // the server security can be improved by checking clients ids from some database of clients
				if errM != nil {
					return errors.WithStack(errM)
				}
				log.Printf("time spent on generating the stamp: %+v", time.Since(t1))
				reply, errR := common.ReqWisdom(servAddr, stamp)
				if errR != nil {
					return errors.WithStack(errR)
				}
				log.Println("reply from server=", reply)
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
