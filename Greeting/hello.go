package Greeting

import "fmt"

func Hello(textMessage string) string {

	if len(textMessage) == 0 {
		return fmt.Sprintf("Message is Empty")
	} else {
		return fmt.Sprintf("Message is %v", textMessage)
	}
}
