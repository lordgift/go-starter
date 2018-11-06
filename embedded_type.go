package main

import (
	"fmt"
)

type Message struct {
	Msg string
}

type SMS struct {
	Message
}

func main() {
	sms := &SMS{
		Message: Message{
			Msg: "Hello, World",
		},
	}

	fmt.Println(sms.Msg)
}
