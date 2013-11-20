all: gozbus

test: gozbus
	./gozbus --server &
	./gozbus

gozbus: gozbus.go zbus/zbus.capnp
	cd zbus; capnp compile -ogo zbus.capnp
	go build -gcflags "-N -l" gozbus.go

clean:
	rm -f *~ zbus/zbus.capnp.go gozbus