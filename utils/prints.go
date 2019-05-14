package utils

import (
	"log"
	"os"
	"runtime"
)

// Err var to add Err message to log
var Err = log.New(os.Stderr,
	"ERROR: ",
	log.Ldate|log.Ltime)

// Error func for print errors
func Error(err error) {
	pc, fn, line, _ := runtime.Caller(1)
	if err != nil {
		Err.Printf("apitruora - in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
	}
}

// PrintMsg func for print message
func PrintMsg(title, a interface{}) {
	log.Printf("apitruora - %v: %#v ", title, a)
}
