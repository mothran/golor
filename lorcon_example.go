package main

/*
#cgo LDFLAGS: -lorcon2
#include <lorcon2/lorcon.h>
#include <lorcon2/lorcon_packet.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
	"time"
)

func main() {

	packet := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")

	channel := C.int(1)
	iface := C.CString("wlan1")

	driver, err := C.lorcon_auto_driver(iface)
	if err != nil {
		panic(err)
	}
	fmt.Println("Name: ", C.GoString(driver.name))

	context, err := C.lorcon_create(iface, driver)
	if err != nil {
		panic(err)
	}

	i, err := C.lorcon_open_injmon(context)
	if err != nil {
		fmt.Println(i)
		fmt.Println("OH GOD WHY!")
		panic(err)
	} else {
		fmt.Println("Monitor mode!")
		C.lorcon_free_driver_list(driver)
	}

	C.lorcon_set_channel(context, channel)

	packNum := 1000
	start := time.Now()

	for i := 0; i < packNum; i++ {
		C.lorcon_send_bytes(context, C.int(len(packet)+1), (*C.u_char)(unsafe.Pointer(&packet[0])))
		// if err != nil{
		// 	break
		// }
		// fmt.Println("Send bytes: ", i)

	}
	elapsed := time.Since(start)
	pps := packNum / (int(elapsed)*1000)
	fmt.Println("pps: ", pps)
	fmt.Println("elaspsed: ", elapsed)

	C.lorcon_close(context)
	C.lorcon_free(context)

}