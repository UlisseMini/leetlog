// +build !windows

package leetlog

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	green  = color.New(color.FgGreen, color.Bold)
	blue   = color.New(color.FgBlue, color.Bold)
	yellow = color.New(color.FgYellow, color.Bold)
	red    = color.New(color.FgRed, color.Bold)
	gray   = color.New(color.FgHiBlack, color.Bold)
)

var DefaultLogger = &Logger{
	Linfo:  log.New(os.Stderr, blue.Sprint(infoPrefix), 0),
	Lwarn:  log.New(os.Stderr, yellow.Sprint(warnPrefix), 0),
	Lerr:   log.New(os.Stderr, red.Sprint(errPrefix), 0),
	Ldebug: log.New(ioutil.Discard, gray.Sprint(debugPrefix), 0),

	OutPrefix: green.Sprint(outPrefix),
}
