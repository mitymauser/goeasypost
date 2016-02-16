package goeasypost

import (
	"log"
	"os"
	"io/ioutil"
)


var ApiKey string = "WCljcleFHNC5j0TzUieXlg"
var ApiVersion string = ""
var HostUri = "https://api.easypost.com/v2"


// SetLogger replaces the default runtime logging configuration with that supplied.
func SetLogger(l *log.Logger) {
	logme = l
}

// DisableLogger sets the logger output to noop
func DisableLogger() {
	SetLogger(log.New(ioutil.Discard,"",0))
}

// SetDefaultLogger sets/restores the internal logger to the default values
func SetDefaultLogger() {
	logme = log.New(os.Stderr,"[GOEASYPOST] ",log.Lshortfile | log.LstdFlags | log.LUTC)
}

func init() {
	SetDefaultLogger()
}

//logme is internal package logger
var logme *log.Logger





