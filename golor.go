package golor

/*
#cgo LDFLAGS: -lorcon2
#include <lorcon2/lorcon.h>
#include <lorcon2/lorcon_packet.h>
*/
import "C"

type Golor interface {
	Context(iface string)
	Send_bytes(data []byte)
	Open_injmon()
	Set_channel(channel int)
}

type golor struct {
}

func list_drivers() lorcon_driver {
	return C.lorcon_list_drivers()
}

func find_driver(driver string) lorcon_driver {
	str := C.CString(driver)
	return C.lorcon_find_driver(str)
}

func auto_driver(iface string) lorcon_driver {
	str := C.CString(iface)
	return C.lorcon_auto_driver(str)
}

func create() lorcon_t {
	return C.lorcon_create(iface, driver)
}

func (c golor) Context(iface string) {
	driver, err := auto_driver(iface)
	if err != nil {
		panic(err)
	}
	context, err := C.lorcon_create(iface, driver)
	if err != nil {
		panic(err)
	}
}