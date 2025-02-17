package main

import (
	"encoding/json"
	"log"
)

func kafkaEventHandler(input []byte) {
	var mailArgs MailArgs
	log.Default().Printf("Data: %s\n", string(input))
	err := json.Unmarshal(input, &mailArgs)
	if err != nil {
		log.Default().Printf("invalid data format: %s\n", err)
		return
	}
	err = sendMail(mailArgs)
	if err != nil {
		log.Default().Printf("invalid data format: %s\n", err)
		return
	}
}
