package bind

import (
	"eighty/log"
	"fmt"
	"net"
	"os"
	"strconv"
)

var einhornFdCount int
const ackMsg = `{"command":"worker:ack","pid":%d}` + "\n"

func envInt(varName string) (int, error) {
	return strconv.Atoi(os.Getenv(varName))
}

func einhornFdMap(n int) int {
	name := fmt.Sprintf("EINHORN_FD_%d", n)
	fd, err := envInt(name)
	if err != nil {
		log.Fatalf("missing %s in environment", name)
	}
	return fd
}

func einhornListen(efd int) (net.Listener, error) {
	fd := einhornFdMap(efd)
	fp := os.NewFile(uintptr(fd), fmt.Sprintf("einhorn@%d", efd))
	defer fp.Close()
	return net.FileListener(fp)
}

func einhornAck() {
	log.Infof("ack'ing to einhorn")
	ctl, err := net.Dial("unix", os.Getenv("EINHORN_SOCK_PATH"))
	if err != nil {
		log.Fatalf("unable to ack to einhorn: %v", err)
	}
	defer ctl.Close()
	_, err = fmt.Fprintf(ctl, ackMsg, os.Getpid())
	if err != nil {
		log.Fatalf("unable to ack to einhorn: %v", err)
	}
}
