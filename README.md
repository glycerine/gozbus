gozbus
======

gozbus: a messaging system in golang.
        Hey hey hey, "Catch Z Bus!"

to install:

    go get github.com/glycerine/gozbus


description
-----------

Gozbus is based on the nanocap transport, our term
for a combination of the nanomsg[1] and Cap'n Proto[2]
technologies.

[1] nanomsg: http://nanomsg.org/
[2] Cap'n Protoc: http://kentonv.github.io/capnproto/


build notes
-----------

to debug with gdb, build like so:

    go build -gcflags "-N -l" gobus.go

this turns of inlining and registerization, so
you can inspect local variables easier.
