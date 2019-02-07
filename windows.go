// +build windows

package leetlog

import (
	"log"
	"os"
)

var defaultLogger = &logger{
	info:  log.New(os.Stderr, infoPrefix, 0),
	warn:  log.New(os.Stderr, warnPrefix, 0),
	err:   log.New(os.Stderr, errPrefix, 0),
	debug: log.New(os.Stderr, debugPrefix, 0),

	outPrefix: outPrefix,
}
