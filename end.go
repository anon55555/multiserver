package multiserver

import (
	"log"
	"os"
	"time"
)

func End(crash, reconnect bool) {
	log.Print("Ending")

	l := GetListener()

	data := make([]byte, 7)
	data[0] = uint8(0x00)
	data[1] = uint8(ToClientAccessDenied)
	if crash {
		data[2] = uint8(AccessDeniedCrash)
	} else {
		data[2] = uint8(AccessDeniedShutdown)
	}
	data[3] = uint8(0x00)
	data[4] = uint8(0x00)
	if reconnect {
		data[5] = uint8(0x01)
	} else {
		data[5] = uint8(0x00)
	}
	data[6] = uint8(0x00)

	for _, clt := range l.addr2peer {
		_, err := clt.Send(Pkt{Data: data})
		if err != nil {
			log.Print(err)
		}

		clt.SendDisco(0, true)
		clt.Close()
	}

	time.Sleep(time.Second)

	if crash {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
