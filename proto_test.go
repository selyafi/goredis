package main

import (
	"fmt"
	"testing"
)

func TestProtocol(t *testing.T) {
	raw := "*3\r\n$3\r\n$3\r\nfoo\r\n$3\r\n\bar\r\n"
	cmd, err := parseCommand(raw)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(cmd)
}
