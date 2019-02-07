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

type Logger struct {
	Linfo  *log.Logger // info level logs
	Lwarn  *log.Logger // warning level logs
	Lerr   *log.Logger // error level logs
	Ldebug *log.Logger // debug level logger

	// output logs for the user
	OutPrefix string
}

// Infional logging
func (l *Logger) Infof(f string, a ...interface{}) { l.Linfo.Printf(f, a...) }
func (l *Logger) Info(f string, a ...interface{})  { l.Linfo.Printf(f, a...) }

// Output (for the user)
func (l *Logger) Printf(f string, a ...interface{}) {
	fmt.Print(l.OutPrefix)
	fmt.Printf(f, a...)
}
func (l *Logger) Print(a ...interface{}) {
	fmt.Print(l.OutPrefix)
	fmt.Print(a...)
}

// Error logging
func (l *Logger) Errorf(f string, a ...interface{}) { l.Lerr.Printf(f, a...) }
func (l *Logger) Error(a ...interface{})            { l.Lerr.Print(a...) }

// Fatal is the same as a call to Error then os.Exit(1)
func (l *Logger) Fatal(a ...interface{}) {
	Error(a...)
	os.Exit(1)
}

// Fatalf is the same as a call to Error then os.Exit(1)
func (l *Logger) Fatalf(f string, a ...interface{}) {
	Errorf(f, a...)
	os.Exit(1)
}

// Default logger methods

// Warning level logs
func (l *Logger) Warnf(f string, a ...interface{}) { l.Lwarn.Printf(f, a...) }
func (l *Logger) Warn(a ...interface{})            { l.Lwarn.Print(a...) }

// Debug level logging
func Debugf(f string, a ...interface{}) { DefaultLogger.Ldebug.Printf(f, a...) }
func Debug(a ...interface{})            { DefaultLogger.Ldebug.Print(a...) }

// Infional logging
func Infof(f string, a ...interface{}) { DefaultLogger.Linfo.Printf(f, a...) }
func Info(a ...interface{})            { DefaultLogger.Linfo.Print(a...) }

// Output (for the user)
func Printf(f string, a ...interface{}) { DefaultLogger.Printf(f, a...) }
func Print(a ...interface{})            { DefaultLogger.Print(a...) }

// Error logging
func Errorf(f string, a ...interface{}) { DefaultLogger.Lerr.Printf(f, a...) }
func Error(a ...interface{})            { DefaultLogger.Lerr.Print(a...) }

// Fatal is the same as a call to Error then os.Exit(1)
func Fatal(a ...interface{})            { DefaultLogger.Fatal(a...) }
func Fatalf(f string, a ...interface{}) { DefaultLogger.Fatalf(f, a...) }

// Warning level logs
func Warnf(f string, a ...interface{}) { DefaultLogger.Lwarn.Printf(f, a...) }
func Warn(a ...interface{})            { DefaultLogger.Lwarn.Print(a...) }
