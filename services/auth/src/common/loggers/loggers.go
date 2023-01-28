package loggers

import "log"

var ErrorLogger = log.New(log.Writer(), "💥: ", log.Ldate|log.Ltime)
var WarningLogger = log.New(log.Writer(), "⚠: ", log.Ldate|log.Ltime)
var InfoLogger = log.New(log.Writer(), "💡: ", log.Ldate|log.Ltime)
