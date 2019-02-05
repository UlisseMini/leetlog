package leetlog

import (
	log "log"
)

var (
	outPrefix  = "[+] "
	infoPrefix = "[~] "
	warnPrefix = "[!] "
	errPrefix  = "[-] "
)

// Informational logging
func Infof(format string, a ...interface{}) {
	defaultLogger.info.Printf(format, a...)
}

// Informational logging
func Info(a ...interface{}) {
	defaultLogger.info.Print(a...)
}

// Output (for the user)
func Printf(format string, a ...interface{}) {
	defaultLogger.out.Printf(format, a...)
}

// Output (for the user)
func Print(a ...interface{}) {
	defaultLogger.out.Print(a...)
}

// Error logging
func Errorf(format string, a ...interface{}) {
	defaultLogger.err.Printf(format, a...)
}

// Error logging
func Error(a ...interface{}) {
	defaultLogger.err.Print(a...)
}

// Warning level logs
func Warnf(format string, a ...interface{}) {
	defaultLogger.warn.Printf(format, a...)
}

// Warning level logs
func Warn(a ...interface{}) {
	defaultLogger.warn.Print(a...)
}

type logger struct {
	out  *log.Logger // output for the user (level 1)
	info *log.Logger // informational logging (level 2)
	warn *log.Logger // warning level logs (level 3)
	err  *log.Logger // error level logs (level 4)

	// log levels to print
	level uint8
}

// Informational logging
func (l *logger) Infof(format string, a ...interface{}) {
	l.info.Printf(format, a...)
}

// Informational logging
func (l *logger) Info(format string, a ...interface{}) {
	l.info.Printf(format, a...)
}

// Output (for the user)
func (l *logger) Printf(format string, a ...interface{}) {
	l.out.Printf(format, a...)
}

// Output (for the user)
func (l *logger) Print(a ...interface{}) {
	l.out.Print(a...)
}

// Error logging
func (l *logger) Errorf(format string, a ...interface{}) {
	l.err.Printf(format, a...)
}

// Error logging
func (l *logger) Error(a ...interface{}) {
	l.err.Print(a...)
}

// Warning level logs
func (l *logger) Warnf(format string, a ...interface{}) {
	l.warn.Printf(format, a...)
}

// Warning level logs
func (l *logger) Warn(a ...interface{}) {
	l.warn.Print(a...)
}
