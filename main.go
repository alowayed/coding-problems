package main

import (
	"fmt"
	"log"
	"time"

	"github.com/alowayed/coding-problems/orth"
)

func run() error {

	dimensions := []int{15, 10}

	o, err := orth.New(dimensions)
	if err != nil {
		return err
	}
	complete, err := o.BridgeComplete()
	if err != nil {
		return err
	}
	for !complete {
		_, err = o.BuildRandom()
		if err != nil {
			return fmt.Errorf("failed to build bridge: %w", err)
		}
		log.Printf("\n%+v", o)
		log.Println("----------")
		complete, err = o.BridgeComplete()
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 300)
	}

	log.Println("--- BRIDGE COMPLETED")

	return nil
}

func main() {
	log.Println("--- Starting")

	if err := run(); err != nil {
		log.Printf("exit: %v", err)
	}

	log.Println("--- End")
}
