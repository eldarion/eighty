package log

import (
	"fmt"
	"log"
	"os"
)

var logger = log.New(os.Stderr, "[eighty] ", log.LstdFlags)

func Printf(level string, format string, v ...interface{}) {
	logger.Printf(fmt.Sprintf("%-9s | %s", level, format), v...)
}

func Infof(format string, v ...interface{}) {
	Printf("INFO", format, v...)
}
