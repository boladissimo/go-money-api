package util

import (
	"log"
	"os"
)

var info *log.Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
var warning *log.Logger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime)
var error *log.Logger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)

// LogInfo log in the info level
func LogInfo(v ...interface{}) {
	info.Println(v...)
}

// LogWarning log in the warning level
func LogWarning(v ...interface{}) {
	warning.Println(v...)
}

// LogError log in the error level
func LogError(v ...interface{}) {
	error.Println(v...)
}
