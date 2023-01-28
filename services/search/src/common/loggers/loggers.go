package loggers

import "log"

var ErrorLogger = log.New(log.Default().Writer(), "SEARCH_SERVICE [ERROR]:", log.Ldate|log.Ltime)
var WarningLogger = log.New(log.Default().Writer(), "SEARCH_SERVICE [WARN]:", log.Ldate|log.Ltime)
var InfoLogger = log.New(log.Default().Writer(), "SEARCH_SERVICE [INFO]:", log.Ldate|log.Ltime)
