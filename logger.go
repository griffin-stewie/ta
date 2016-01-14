package main

import (
	"github.com/hashicorp/logutils"
	"log"
	"os"
)

// LogLevel is type for page type
type LogLevel int

func (s LogLevel) String() string {
	switch s {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warn:
		return "WARN"
	case Error:
		return "ERROR"
	default:
		return "ERROR"
	}
}

// LogLevel enum
const (
	Debug LogLevel = iota + 1
	Info
	Warn
	Error
)

// SetupLogger is
func SetupLogger(level LogLevel) {
	var loglevel = logutils.LogLevel(level.String())
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: loglevel,
		Writer:   os.Stderr,
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(filter)
}
