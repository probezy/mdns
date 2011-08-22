package zeroconf

import (
	dns "github.com/miekg/godns"
)

var (
        Local = NewZone("local.")
	Listener = Listen(Local) // start mcast listner
)

func init() {
        go Local.mainloop()
}

func Publish(rr dns.RR) {
        Local.Add(&Entry{
                expires: 2 ^ 31, // never
                publish: true,
                rr:      rr,
        })
}

