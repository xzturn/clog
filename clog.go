// xzturn package: colored console log
// 
// Author: tengming.sp
//
package clog

import (
    "log"
    "os"
    "github.com/fatih/color"
)

////////////////////////////////////////////////////////////////////////////////

type ConsoleLogger struct {
    *log.Logger
}

func NewConsoleLogger(prefix string) *ConsoleLogger {
    return &ConsoleLogger{ log.New(os.Stdout, prefix, log.LstdFlags) }
}

func (c *ConsoleLogger) Error(format string, v ...interface{}) {
    color.Set(color.FgRed)
    c.Printf(format, v...)
    color.Unset()
}

func (c *ConsoleLogger) Success(format string, v ...interface{}) {
    color.Set(color.FgGreen)
    c.Printf(format, v...)
    color.Unset()
}
