package bind

import (
	"eighty/log"
	"github.com/stripe/go-einhorn/einhorn"
	"net"
)

func einhornListen(efd int) (net.Listener, error) {
	return einhorn.GetListener(efd)
}

func einhornAck() {
	log.Infof("ack'ing to einhorn")
	einhorn.Ack()
}
