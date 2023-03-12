package main

import (
	"log"
	"os"
	"testing"
)

const testPort = 9002

func TestMain(t *testing.M) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	go func(){
		if err := tcpServe(testPort); err != nil {
			log.Printf("error: %+v", err)
			os.Exit(1)
		}
	}()
	os.Exit(t.Run())
}
