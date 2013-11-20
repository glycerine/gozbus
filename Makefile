all: gozbus

gozbus:
	capnp compile -ogo zbus.capnp
	go build -gcflags "-N -l" gozbus.go

clean:
	rm -f *~ zbus.capnp.go gozbus