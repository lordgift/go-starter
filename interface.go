package main

import (
	"fmt"
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

func main() {
	var sender Sender
	sender = &SMS{"0000000000"}
	sender.Send("Happy New Year!!")

	sender = &Email{"email@email.com"}
	sender.Send("Happy New Year!!")
}
