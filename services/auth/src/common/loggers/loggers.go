package loggers

import "log"

var (
	ErrorLogger   *log.Logger
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
)

func InitLoggers() {
	ErrorLogger = log.New(log.Writer(), "💥: ", log.Ldate|log.Ltime)
	WarningLogger = log.New(log.Writer(), "⚠: ", log.Ldate|log.Ltime)
	InfoLogger = log.New(log.Writer(), "💡: ", log.Ldate|log.Ltime)
}
