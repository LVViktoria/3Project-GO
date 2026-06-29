package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {

	host := "smtp.gmail.com"
	port := "25"

	fmt.Println("Connecting...")

	client, err := smtp.Dial(host + ":" + port)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	fmt.Println("Connected successfully!")
}
