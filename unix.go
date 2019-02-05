// +build !windows

package leetlog

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	green  = color.New(color.FgGreen, color.Bold)
	blue   = color.New(color.FgBlue, color.Bold)
	yellow = color.New(color.FgYellow, color.Bold)
	red    = color.New(color.FgRed, color.Bold)
)

var defaultLogger = &logger{
	out:  log.New(os.Stdout, green.Sprint(outPrefix), 0),
	info: log.New(os.Stderr, blue.Sprint(infoPrefix), 0),
	warn: log.New(os.Stderr, yellow.Sprint(warnPrefix), 0),
	err:  log.New(os.Stderr, red.Sprint(errPrefix), 0),
}

// Return a bold version of s
func bold(s string) string {
	return fmt.Sprintf("\033[1m%s\033[0m", s)
}
