package bind

import (
	"eighty/log"
	"net"
	"strconv"
	"strings"
)

var bindMode int

func Listen(bind string) (net.Listener, error) {
	if strings.HasPrefix(bind, "einhorn") {
		bits := strings.Split(bind, "@")
		efd, err := strconv.Atoi(bits[1])
		if err != nil {
			log.Fatalf("%v", err)
		}
		bindMode = 0x02
		return einhornListen(efd)
	}
	bindMode = 0x01
	return netListen(bind)
}

func Ready() {
	if bindMode == 0x02 {
		einhornAck()
	}
}
