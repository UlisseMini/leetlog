// +build windows

package leetlog

import (
	"io/ioutil"
	"log"
	"os"
)

var DefaultLogger = &Logger{
	Linfo:  log.New(os.Stderr, infoPrefix, 0),
	Lwarn:  log.New(os.Stderr, warnPrefix, 0),
	Lerr:   log.New(os.Stderr, errPrefix, 0),
	Ldebug: log.New(ioutil.Discard, debugPrefix, 0),

	OutPrefix: outPrefix,
}
