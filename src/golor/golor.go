package golor

/*
#cgo LDFLAGS: -lorcon2
#include <lorcon2/lorcon.h>
#include <lorcon2/lorcon_packet.h>
*/
import "C"

import (
	"unsafe"
)


type Golor struct {
	iface 	*C.char
	channel C.int
	driver 	*_Ctype_lorcon_driver_t
	context *C.struct_lorcon
}

// private methods

func (golor *Golor) auto_driver() {
	driver, err := C.lorcon_auto_driver(golor.iface)
	if err != nil {
		panic(err)
	}
	golor.driver = driver
}

func (golor *Golor) create() {
	context, err := C.lorcon_create(golor.iface, golor.driver)
	if err != nil {
		panic(err)
	}
	golor.context = context
}


// public methods 

func (golor *Golor) Open_injmon() {
	_, err := C.lorcon_open_injmon(golor.context)
	if err != nil {
		// if you error here it is most likely because you need to run
		// as root for injection.
		panic(err)
	}
}

func (golor *Golor) Set_chan(channel int) {
	golor.channel = C.int(channel)

	C.lorcon_set_channel(golor.context, golor.channel)
}

func (golor *Golor) Chan() int {
	return int(golor.channel)
}

func (golor *Golor) Send_bytes(packet []byte) int {
	// will return the number of bytes written, this might include the radiotap outbound header.
	var sent = C.lorcon_send_bytes(golor.context, C.int(len(packet)+1), (*C.u_char)(unsafe.Pointer(&packet[0])))
	return int(sent)
} 

func (golor *Golor) Context(iface string) {
	
	// assign the interface
	golor.iface = C.CString(iface)

	// grab the driver info
	golor.auto_driver()

	// create the context
	golor.create()

}

func (golor *Golor) Getdriver() string {
	return C.GoString(golor.driver.name)
}

func (golor *Golor) Version() int {
	return int(C.lorcon_get_version())
}

func (golor *Golor) Close() {
	C.lorcon_free_driver_list(golor.driver)
	C.lorcon_close(golor.context)
}