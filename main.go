package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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

func main() {
	url, exists := os.LookupEnv("RETRO_RACH_URL")
	if !exists {
		log.Fatal("'RETRO_RACH_URL' environmental variable is undefined")
	}

	sendMessage(url, "oh hellooooooo there")
}
