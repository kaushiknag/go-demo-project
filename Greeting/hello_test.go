package Greeting_test

import (
	"github.com/kaushiknag/codebuild_go_demo/Greeting"
	"testing"
)

func TestHello(t *testing.T) {

	emptyResult := Greeting.Hello("")

	if emptyResult != "Message is Empty" {
		t.Errorf("hello function failed with error %v", emptyResult)
	} else {
		t.Logf("hello function success with expected message as %v", emptyResult)
	}

	result := Greeting.Hello("HelloWorld")

	if result != "Message is HelloWorld" {
		t.Errorf("hello function failed with error %v", result)
	} else {
		t.Logf("hello function success with expected message as %v", result)
	}
}

func TestHelloEmptyArg(t *testing.T) {

	emptyResult := Greeting.Hello("")

	if emptyResult != "Message is Empty" {
		t.Errorf("hello function failed with error %v", emptyResult)
	} else {
		t.Logf("hello function success with expected message as %v", emptyResult)
	}
}

func TestHelloValidArg(t *testing.T) {

	result := Greeting.Hello("HelloWorld")

	if result != "Message is HelloWorld" {
		t.Errorf("hello function failed with error %v", result)
	} else {
		t.Logf("hello function success with expected message as %v", result)
	}
}
