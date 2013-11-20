package main

import (
	"testing"
	nn "github.com/op/go-nanomsg"
    "os/exec"
	"strings"
	cv "github.com/smartystreets/goconvey/convey"
)


func TestDemo(t *testing.T) {
    t.Logf("use the -test.v flag to see this output")
}


// basic functionality:


// starting server should bind the supplied endpoint
func TestServerBinds(t *testing.T) {
	ServerBindHelper(t, ZBUS_ADDR, ZBUS_ADDR)
}

// and the netstat validation should fail if we
// give the wrong endpoint:

func TestBadEndpointMeansServerEndpointTestShouldFail(t *testing.T) {
	cv.ShouldPanic(func() { panic("test the goconvey ShouldPanic function") })
    cv.ShouldPanic(func() { ServerBindHelper(t, "tcp://127.0.0.1:1776", "tcp://127.0.0.1:1779")})	 
    cv.ShouldPanic(func() { ServerBindHelper(t, "tcp://127.0.0.1:1777", "tcp://127.0.0.1:1778")})
    cv.ShouldNotPanic(func() { ServerBindHelper(t, "tcp://127.0.0.1:1779", "tcp://127.0.0.1:1779")})
}


// make addr separate from ZBUS_ADDR, so we
// can validate that the test detects a problem
// when they are different.
func ServerBindHelper(t *testing.T, addr_use string, addr_expect string) {
	nnzbus, err := nn.NewSocket(nn.AF_SP, nn.PAIR)
	if err != nil { t.Fatal(err) }
	defer nnzbus.Close()

	startZBus(nnzbus, addr_use)

	// To test the test, for example, expect failure if we gave a wrong address here:
    // startZBus(nnzbus, "tcp://127.0.0.1:1777")
	
	out, err := exec.Command("netstat", "-nuptl").Output()
	if err != nil { t.Fatal(err) }
	
	lines := strings.Split(string(out), "\n")	
	needle := strings.SplitAfter(addr_expect, "//")[1]
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

	if !found { 
		t.Logf("gozbus server was not listening on %v as expected", addr_expect) 
		panic("no gozbus server at expected endpoint")
	}
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
