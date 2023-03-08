package main

import (
	"github.com/pkg/errors"
	"log"

	hc "github.com/catalinc/hashcash"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := func() error {
		h := hc.NewStd() // or .New(bits, saltLength, extra)

		// Mint a new stamp
		stamp, err := h.Mint("something")
		if err != nil {
			return errors.WithStack(err)
		}
		log.Printf("stamp: %+v", stamp)

		// Check a stamp
		valid := h.Check("1:20:161203:something::+YO19qNZKRs=:a31a2")
		if valid {
			log.Println("Valid")
		} else {
			log.Println("Invalid")
		}
		return nil
	}(); err != nil {
		log.Printf("err: %+v", err)
	}
	log.Println("finished")
}