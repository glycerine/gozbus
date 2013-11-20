package main

import (
	//"os"
	"testing"
	nn "github.com/op/go-nanomsg"
    "os/exec"
	//"regexp"
	"strings"
	cv "github.com/smartystreets/goconvey/convey"
)


func TestDemo(t *testing.T) {
    t.Logf("use the -test.v flag to see this output")
}


// basic functionality:


// starting server should bind the supplied endpoint
func TestServerBinds(t *testing.T) {
	ServerBindHelper(t, ZBUS_ADDR)
}

// and the netstat validation should fail if we
// give the wrong endpoint:

func badEndpointMeansServerEndpointTestShouldFail(t *testing.T) {
    ServerBindHelper(t, "tcp://127.0.0.1:99992423423")
}


// make addr separate from ZBUS_ADDR, so we
// can validate that the test detects a problem
// when they are different.
func ServerBindHelper(t *testing.T, addr string) {
	nnzbus, err := nn.NewSocket(nn.AF_SP, nn.PAIR)
	if err != nil { t.Fatal(err) }
	
	startZBus(nnzbus, addr)
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
			t.Logf("'netstat -nuptl | grep %v | grep LISTEN' found: %v", needle, haystack) 
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


// go convey try out:

func TestSpec(t *testing.T) {
    var x int

    // Only pass t into top-level Convey calls
    cv.Convey("Given some integer with a starting value", t, func() {
        x = 1

        cv.Convey("When the integer is incremented", func() {
            x++

            cv.Convey("The value should be greater by one", func() {
                cv.So(x, cv.ShouldEqual, 2)
            })
        })
    })
}
