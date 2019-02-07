package leetlog

import (
	"fmt"
	log "log"
	"os"
)

var (
	outPrefix   = "[+] "
	infoPrefix  = "[*] "
	warnPrefix  = "[!] "
	errPrefix   = "[-] "
	debugPrefix = "[~] "
)

type logger struct {
	info  *log.Logger // info level logs
	warn  *log.Logger // warning level logs
	err   *log.Logger // error level logs
	debug *log.Logger // debug level logger

	// output logs for the user
	outPrefix string
}

// Infional logging
func (l *logger) Infof(f string, a ...interface{}) { l.info.Printf(f, a...) }
func (l *logger) Info(f string, a ...interface{})  { l.info.Printf(f, a...) }

// Output (for the user)
func (l *logger) Printf(f string, a ...interface{}) {
	fmt.Print(l.outPrefix)
	fmt.Printf(f, a...)
}
func (l *logger) Print(a ...interface{}) {
	fmt.Print(l.outPrefix)
	fmt.Print(a...)
}

// Error logging
func (l *logger) Errorf(f string, a ...interface{}) { l.err.Printf(f, a...) }
func (l *logger) Error(a ...interface{})            { l.err.Print(a...) }

// Fatal is the same as a call to Error then os.Exit(1)
func (l *logger) Fatal(a ...interface{}) {
	Error(a...)
	os.Exit(1)
}

// Fatalf is the same as a call to Error then os.Exit(1)
func (l *logger) Fatalf(f string, a ...interface{}) {
	Errorf(f, a...)
	os.Exit(1)
}

// Default logger methods

// Warning level logs
func (l *logger) Warnf(f string, a ...interface{}) { l.warn.Printf(f, a...) }
func (l *logger) Warn(a ...interface{})            { l.warn.Print(a...) }

// Debug level logging
func Debugf(f string, a ...interface{}) { defaultLogger.debug.Printf(f, a...) }
func Debug(a ...interface{})            { defaultLogger.debug.Print(a...) }

// Infional logging
func Infof(f string, a ...interface{}) { defaultLogger.info.Printf(f, a...) }
func Info(a ...interface{})            { defaultLogger.info.Print(a...) }

// Output (for the user)
func Printf(f string, a ...interface{}) { defaultLogger.Printf(f, a...) }
func Print(a ...interface{})            { defaultLogger.Print(a...) }

// Error logging
func Errorf(f string, a ...interface{}) { defaultLogger.err.Printf(f, a...) }
func Error(a ...interface{})            { defaultLogger.err.Print(a...) }

// Fatal is the same as a call to Error then os.Exit(1)
func Fatal(a ...interface{})            { defaultLogger.Fatal(a...) }
func Fatalf(f string, a ...interface{}) { defaultLogger.Fatalf(f, a...) }

// Warning level logs
func Warnf(f string, a ...interface{}) { defaultLogger.warn.Printf(f, a...) }
func Warn(a ...interface{})            { defaultLogger.warn.Print(a...) }
