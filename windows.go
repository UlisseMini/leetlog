// +build windows

package leetlog

import (
	"log"
	"os"
)

var defaultLogger = &logger{
	out:  log.New(os.Stdout, outPrefix, 0),
	info: log.New(os.Stderr, infoPrefix, 0),
	warn: log.New(os.Stderr, warnPrefix, 0),
	err:  log.New(os.Stderr, errPrefix, 0),
}
