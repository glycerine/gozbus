// copyright(c) 2013, the gozbus authors.
// license: MIT
//
// gozbus: a messaging system in golang.
//         Hey hey hey, "Catch Z Bus!"
//

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	//"time"
	capn "github.com/glycerine/go-capnproto"
	"github.com/glycerine/gozbus/zbus"
	nn "github.com/op/go-nanomsg"

	// for parsing command line options:
	// https://github.com/spf13/cobra
)

const ZBUS_ADDR = "tcp://127.0.0.1:1776"

func startZBus(nnzbus *nn.Socket, addr string) {

	_, err := nnzbus.Bind(addr)
	if err != nil {
		fmt.Printf("could not bind addr '%s': %v", addr, err)
		panic(err)
	}
	fmt.Printf("[pid %d] gozbus server: startZbus bound endpoint '%s'.\n", os.Getpid(), addr)
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
	msg := "hello from pid " + strconv.Itoa(pid)

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
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(myMsg)
	capMsg, err := capn.ReadFromStream(buf, nil)
	if err != nil {
		log.Fatal(err)
	}

	z := gozbus.ReadRootZ(capMsg)
	mydate := z.Zdate()

	fmt.Printf("[pid %d] recvZDate got ZDate message: year %d, month %d, day %d\n",
		os.Getpid(), mydate.Year(), mydate.Month(), mydate.Day())
}

func main() {
	var err error
	var nnzbus *nn.Socket
	pid := os.Getpid()

	var isServer bool = false
	if len(os.Args) > 1 && os.Args[1] == "--server" {
		isServer = true
	}

	if isServer {
		// server code, binds the bus to start it.
		nnzbus, err = nn.NewSocket(nn.AF_SP, nn.REP)
		if err != nil {
			log.Fatal(err)
		}
		defer nnzbus.Close()

		startZBus(nnzbus, ZBUS_ADDR)
		recvMsgOnZBus(nnzbus)
		sendZDate(nnzbus)

	} else {
		// client code, connects to the bus.
		nnzbus, err = nn.NewSocket(nn.AF_SP, nn.REQ)
		if err != nil {
			log.Fatal(err)
		}
		defer nnzbus.Close()

		if _, err = nnzbus.Connect(ZBUS_ADDR); err != nil {
			log.Fatal(err)
		}

		sayHello(nnzbus)
		recvZDate(nnzbus)
	}

	fmt.Printf("[pid %d] done.\n", pid)
}
