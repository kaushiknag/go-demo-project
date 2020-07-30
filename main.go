package main

import (
	"fmt"
	"github.com/kaushiknag/go-demo-project/Greeting"
)

func main() {

	greetMessageEmpty := Greeting.Hello("")
	fmt.Println(greetMessageEmpty)

	greetMessage := Greeting.Hello("HelloWorld")
	fmt.Println(greetMessage)
}
