package main

import (
	"fmt"
	"log"
)

type Sender interface {
	Send(msg string) error
}
type SMS struct {
	phoneNo string
}

type Email struct {
	destEmail string
}

func (sms *SMS) Send(msg string) error {
	fmt.Printf("Sending SMS To %s\n%s\n\n", sms.phoneNo, msg)
	return nil
}

func (email *Email) Send(msg string) error {
	fmt.Printf("Sending Email To %s\n%s\n\n", email.destEmail, msg)
	return nil
}

func SendMessage(sender Sender, message string) {
	log.Println("====== Sending ======")
	sender.Send(message)
	log.Println("====== Sent ======")
}

func main() {
	SendMessage(&SMS{"0000000000"}, "Happy New Year!!")
	SendMessage(&Email{"email@mail.com"}, "Happy New Year!!")
}
