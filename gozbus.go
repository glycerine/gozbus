// copyright(c) 2013, the gobus authors.
// license: MIT or Apache2, your choice.
// gobus: a nanocap based messaging system in golang
//

// to debug with gdb, build like so:
//    go build -gcflags "-N -l" gobus.go

package main

import (
	//"bytes"
	"fmt"
	nn "github.com/op/go-nanomsg"
	"log"
	"os"
	"strconv"
	//"time"
)

const BUS_ADDR = "tcp://127.0.0.1:1776"


func startBus(nnbus *nn.Socket) {

    _, err := nnbus.Bind(BUS_ADDR)
	if err != nil {
		fmt.Printf("bus already started?; proceeding.\n")
	}
	fmt.Printf("busclient: startBus bound endpoint '%s'.\n", BUS_ADDR)
}

func recvMsgOnBus(nnbus *nn.Socket) {
	
	// receive, synchronously so flags == 0
	var flags int = 0
	heardBuf, err := nnbus.Recv(flags)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("I heard: '%s'.\n", heardBuf)
}

func sayHello(nnbus *nn.Socket) {
	whoami := os.Getpid()
	msg    := "hello from pid " + strconv.Itoa(whoami)
	
	nnbus.Send([]byte(msg), 0)
	fmt.Printf("busclient: sayHello sent msg '%s'.\n", msg)
}


func main() {
	var err error
	var nnbus *nn.Socket

	nnbus, err = nn.NewSocket(nn.AF_SP, nn.PAIR)
	if err != nil { log.Fatal(err) }

	var isServer bool = false
	if len(os.Args) > 1 && os.Args[1] == "--server" {
		isServer = true
	}

	if isServer {
		// server code, binds the bus to start it.
		startBus(nnbus)
		recvMsgOnBus(nnbus)

	} else {
		// client code, connects to the bus.
		if _, err = nnbus.Connect(BUS_ADDR); err != nil {
			log.Fatal(err)
		}

		sayHello(nnbus)
	}

	// wait 2 minutes
	fmt.Printf("done.\n")
}
