/*
 * Copyright (c) 2016 Stewart Buskirk <mitymauser@gmail.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */
package goeasypost

import (
	"io/ioutil"
	"log"
	"os"
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
	SetLogger(log.New(ioutil.Discard, "", 0))
}

// SetDefaultLogger sets/restores the internal logger to the default values
func SetDefaultLogger() {
	logme = log.New(os.Stderr, "[GOEASYPOST] ", log.Lshortfile|log.LstdFlags|log.LUTC)
}

func init() {
	SetDefaultLogger()
}

//logme is internal package logger
var logme *log.Logger
