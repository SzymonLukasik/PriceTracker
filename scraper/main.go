package main

import (
	"fmt"
	"pricetracker/scraper/internal/scheduler"
	// FIXME ADD CLI"github.com/alecthomas/kingpin"
)

func main() {
	fmt.Println("hello")
	duration := "2s"
	scheduler.Schedule(duration)
	fmt.Println("XD")
}
