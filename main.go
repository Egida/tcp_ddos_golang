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