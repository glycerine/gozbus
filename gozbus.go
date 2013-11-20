// copyright(c) 2013, the gozbus authors.
// license: MIT or Apache2, your choice.
//
// gozbus: a messaging system in golang.
//         As a European might say, "Catch Z Bus!"
//
// Gozbus is based on the nanocap transport, our term
// for a combination of the nanomsg[1] and Cap'n Proto[2]
// technologies.
//
// [1] nanomsg: http://nanomsg.org/
// [2] Cap'n Protoc: http://kentonv.github.io/capnproto/

//
// build notes:
// to debug with gdb, build like so:
//    go build -gcflags "-N -l" gobus.go
//  this turns of inlining and registerization, so
//  you can inspect local variables easier.

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	//"time"
	nn "github.com/op/go-nanomsg"
	"github.com/glycerine/gozbus/zbus"
	capn "github.com/jmckaskill/go-capnproto"
)

const ZBUS_ADDR = "tcp://127.0.0.1:1776"


func startZBus(nnzbus *nn.Socket) {

    _, err := nnzbus.Bind(ZBUS_ADDR)
	if err != nil {
		fmt.Printf("zbus already started?; proceeding.\n")
	}
	fmt.Printf("[pid %d] gozbus server: startZbus bound endpoint '%s'.\n", os.Getpid(), ZBUS_ADDR)
}

func recvMsgOnZBus(nnzbus *nn.Socket) {
	pid := os.Getpid()
	
	// receive, synchronously so flags == 0
	var flags int = 0
	heardBuf, err := nnzbus.Recv(flags)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[pid %d] gozbus server: I heard: '%s'.\n", pid, heardBuf)
}

func sayHello(nnzbus *nn.Socket) {
	pid := os.Getpid()
	msg    := "hello from pid " + strconv.Itoa(pid)
	
	nnzbus.Send([]byte(msg), 0)
	fmt.Printf("[pid %d] gozbus client: sayHello sent msg '%s'.\n", pid, msg)
}

func sendZDate(nnzbus *nn.Socket) {

    // Create Zdate and Write to nnzbus.
	//
	// Z is our universal type, a discriminated union.
	// Hence Z provides run-time type identification
	// for all message types defined in zbus.capnp
	//
	s := capn.NewBuffer(nil)
	z := gozbus.NewRootZ(s)
	d := gozbus.NewRootZdate(s)
	z.SetZdate(d)

	d.SetYear(2004)
	d.SetMonth(12)
	d.SetDay(7)

	buf := bytes.Buffer{}
	s.WriteTo(&buf)

	nnzbus.Send(buf.Bytes(), 0)

	fmt.Printf("[pid %d] sendZDate sent ZDate message: year %d, month %d, day %d\n", 
		os.Getpid(), d.Year(), d.Month(), d.Day())
}

func recvZDate(nnzbus *nn.Socket) {

    // Read a Z message that is expected to be a Zdate from nnzbus
	myMsg, err := nnzbus.Recv(0)
	if err != nil { log.Fatal(err) }

	buf := bytes.NewBuffer(myMsg)
	capMsg, err := capn.ReadFromStream(buf, nil)
	if err != nil { log.Fatal(err) }

	z := gozbus.ReadRootZ(capMsg)
	mydate := z.Zdate()

	fmt.Printf("[pid %d] recvZDate got ZDate message: year %d, month %d, day %d\n", 
		os.Getpid(), mydate.Year(), mydate.Month(), mydate.Day())
}

func main() {
	var err error
	var nnzbus *nn.Socket
	pid := os.Getpid()

	nnzbus, err = nn.NewSocket(nn.AF_SP, nn.PAIR)
	if err != nil { log.Fatal(err) }

	var isServer bool = false
	if len(os.Args) > 1 && os.Args[1] == "--server" {
		isServer = true
	}

	if isServer {
		// server code, binds the bus to start it.
		startZBus(nnzbus)
		recvMsgOnZBus(nnzbus)
		sendZDate(nnzbus)

	} else {
		// client code, connects to the bus.
		if _, err = nnzbus.Connect(ZBUS_ADDR); err != nil {
			log.Fatal(err)
		}

		sayHello(nnzbus)
		recvZDate(nnzbus)
	}

	// wait 2 minutes
	fmt.Printf("[pid %d] done.\n", pid)
}
