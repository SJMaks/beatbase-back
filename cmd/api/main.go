package main

import (
	"fmt"
	"log"

	"github.com/sjmaks/beatbase-back/internal/core/app"
)

func main() {
	if err := safeRun(); err != nil {
		log.Fatalf("failed to start the app: %v", err)
	}
}

func safeRun() (err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("PANIC: %v", r)
			err = fmt.Errorf("PANIC: %v", r)
		}
	}()
	return app.Run()
}
