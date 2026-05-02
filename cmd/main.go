package main

import (
	"log-generator/internal"
)

func main() {
	cfg := internal.LoadConfig()
	internal.StartEmitter(cfg)

	select {} // keep running
}