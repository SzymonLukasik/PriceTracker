package scheduler

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func getData(resourceUrl string) {
	fmt.Printf("fetching data from %s\n", resourceUrl)
	fmt.Println("Data fetched successfully")
}

func Schedule(duration string) {
	s := gocron.NewScheduler(time.UTC)
	s.Every(duration).Do(getData, "www.google.com")
	s.StartBlocking()
}
