package main

import (
	"golor"
	"fmt"
	"os"
)

func usage() {
    fmt.Fprintf(os.Stderr, "usage: %s [interface]\n", os.Args[0])
    os.Exit(2)
}

func main() {
	if len(os.Args) <= 1 {
		usage()
	}
	var iface = os.Args[1]

	var golor golor.Golor

	packet := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")

	golor.Context(iface)
	fmt.Printf("Current driver: %s\n", golor.Getdriver())
	golor.Open_injmon()
	golor.Set_chan(11)
	
	fmt.Printf("Current chan: %d\n", golor.Chan())
	fmt.Printf("Current vers: %d\n", golor.Version())
	
	var sent = golor.Send_bytes(packet)
	
	fmt.Printf("Done, sent: %d\n", sent)
	golor.Close()
}