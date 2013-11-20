package main

import (
	"testing"
	nn "github.com/op/go-nanomsg"
    "os/exec"
	//"regexp"
	"strings"
)

func TestDemo(t *testing.T) {
    t.Logf("use the -test.v flag to see this output")
}

// basic functionality:


// starting server should bind the supplied endpoint
func TestServerbinds(t *testing.T) {
	nnzbus, err := nn.NewSocket(nn.AF_SP, nn.PAIR)
	if err != nil { t.Fatal(err) }

	startZBus(nnzbus, ZBUS_ADDR)
	// To test the test, for example, expect failure if we gave a wrong address here:
    // startZBus(nnzbus, "tcp://127.0.0.1:1777")

	out, err := exec.Command("netstat", "-nuptl").Output()
	if err != nil { t.Fatal(err) }

	lines := strings.Split(string(out), "\n")

	needle := strings.SplitAfter(ZBUS_ADDR, "//")[1]
	//e.g.	needle := "127.0.0.1:1776"

	t.Logf("using netstat -nuptl to look for %v LISTEN line.\n", needle)

	found := false
	for _, haystack := range lines {
		if strings.Contains(haystack, needle) && 
			strings.Contains(haystack, "LISTEN") { 
			found = true 
			break
		}
	}

	if !found { t.Fatal("gozbus server was not",
		"listening on", ZBUS_ADDR, "as expected.") }
}

// client should be able to send to live server

// client should be receive from live server

// server should be able to send to client

// server should be able to receive from client


