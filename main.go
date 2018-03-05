package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func sendMessage(url, message string) {
	log.Printf("Sending message %s", message)

	client := http.Client{}
	body := strings.NewReader(fmt.Sprintf(`{"text": "%s"}`, message))
	rsp, err := client.Post(url, "application/json", body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Sent message. Response %s", rsp.Status)
}

func everyHour(f func()) {
	d := time.Second * 3
	t := time.NewTimer(d)
	for {
		select {
		case <-t.C:
			f()
			t.Reset(d)
		}
	}
}

func nextMonday(now time.Time) time.Time {
	var dayDiff int
	if now.Weekday() == time.Sunday {
		dayDiff = 7
	} else {
		dayDiff = 7 - int(now.Weekday()-time.Monday)
	}
	oneDay := time.Hour * 24
	return now.Add(oneDay * time.Duration(dayDiff))
}

func main() {
	url, exists := os.LookupEnv("RETRO_RACH_URL")
	if !exists {
		log.Fatal("'RETRO_RACH_URL' environmental variable is undefined")
	}

	now := time.Now().Add(time.Hour * 24)
	nextMonday := nextMonday(now)
	nextNotification := time.Date(nextMonday.Year(), nextMonday.Month(), nextMonday.Day(), 9, 30, 0, 0, nextMonday.Location())
	log.Printf("Now: %s, nextMonday: %s, nextNotifcation: %s", now, nextMonday, nextNotification)

	everyHour(func() {
		fmt.Println("AAA")
		_ = url
	})
}
