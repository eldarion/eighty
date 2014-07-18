package log

import (
	"fmt"
	"log"
	"os"
)

var logger = log.New(os.Stdout, fmt.Sprintf("[eighty %d] ", os.Getpid()), log.LstdFlags)

func Printf(level string, format string, v ...interface{}) {
	logger.Printf(fmt.Sprintf("%-9s | %s", level, format), v...)
}

func Debugf(format string, v ...interface{}) {
	Printf("DEBUG", format, v...)
}

func Infof(format string, v ...interface{}) {
	Printf("INFO", format, v...)
}

func Warnf(format string, v ...interface{}) {
	Printf("WARNING", format, v...)
}

func Fatalf(format string, v ...interface{}) {
	Printf("FATAL", format, v...)
	os.Exit(1)
}
