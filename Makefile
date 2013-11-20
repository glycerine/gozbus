all: gozbus

test: gozbus
	./gozbus --server &
	./gozbus

gozbus: gozbus.go zbus.capnp
	capnp compile -ogo zbus.capnp
	go build -gcflags "-N -l" gozbus.go

clean:
	rm -f *~ zbus.capnp.go gozbus