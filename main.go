package main

import (
	"log"
	"ssdp/ssdp"
)

func main() {
	// start device discovery
	discovery := ssdp.NewDiscovery()

	for device := range discovery.Reporter {
		log.Printf("device: %s\n", device)
	}
}