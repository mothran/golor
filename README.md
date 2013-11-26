#golor

GoLor is a GoLang binding for the Lorcon2 library

To get started try:

	go install golor
	sudo -E -c "go run main.go [INTERFACE]"


The interface is modeled after pylorcon2 (this one: ) and for examples of how to interact with the library, check main.go

#Status
Working, modeled after pylorcon2

Currently it is able to open an interface in inj/monitor mode, set and get the chan, and inject raw bytes. 

## Functional Methods

	Open_injmon()
	Set_chan(channel int)
	Chan() int
	Send_bytes(packet []byte) int
	Context(iface string)
	Getdriver() string
	Version() int
	Close()

## Not implemented
	
	Get or Sets for MAC addr's
	Get or Sets for VAP
	datalink func's
	Anything not mentioned above.
	An error handeling system, currently it panics in the library.