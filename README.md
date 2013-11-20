gozbus
======

gozbus: a messaging system in golang.
        Hey hey hey, "Catch Z Bus!"

to install:

    go get github.com/glycerine/gozbus


description
-----------

Gozbus is a messaging system based on the nanocap transport, 
our term for a combination of the nanomsg[1] and Cap'n Proto[2]
technologies.

Nanomsg is the successor to ZeroMQ, and
provides multiple patterns for in-process, in-host,
and inter-host messaging. The peer-to-peer, publish-subscribe,
enterprise bus, pipeline, and surveyor protcols can be
leveraged for scalable messaging.

Cap'n Proto is the successor to ProtocolBuffers, and 
provides highly efficient encoding
and decoding of messages based on a strongly typed schema
language. Capnp bindings are available for Golang, 
C++, and Python. We use the schema handling portion only,
as the RPC part of Cap'n Proto isn't available at this time.

[1] nanomsg: http://nanomsg.org/

[2] Cap'n Proto: http://kentonv.github.io/capnproto/

These two pre-requisite libraries should be downloaded and installed
prior to building gozbus.

build notes
-----------

to debug with gdb, build like so:

    go build -gcflags "-N -l" gozbus.go

this turns off inlining and registerization, so
you can inspect local variables easier.
