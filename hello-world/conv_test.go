package main

import (
	"testing"
)

func TestConv(t *testing.T) {
	num, err := stringToNum("123456789")
	if err != nil {
		t.Fatal(err)
	}
	if num != 123456789 {
		t.Fatal("Number don't match")
	}
}
func TestFailConv(t *testing.T) {
	_, err := stringToNum("")
	if err == nil {
		t.Fatal(err)
	}
}
