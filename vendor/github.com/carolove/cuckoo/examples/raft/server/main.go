package main

import (
	"os"

	"github.com/carolove/cuckoo"
)

const (
	PORT = ":9090"
)

func main() {
	s, err := cuckoo.NewNode()
	if err != nil {
		os.Exit(-1)
	}

	go func() {
		s.Start()
	}()

}
