package log

import (
    "log"
    "os"

)


func init() {
    f, err := os.OpenFile("./full.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

    if err != nil {
        log.Fatalf("error opening file: %v", err)
    }
    defer f.Close()
    log.SetOutput(f)

}


// InPrint`ifo ...
func Printf(format string, v ...interface{}) {
    log.Printf(format, v...)
}

func Println(format string, v ...interface{}) {

    log.Printf(format, v...)
}

/*
// Warn ...
func Warn(format string, v ...interface{}) {
    log.Warnf(format, v...)
}

// Error ...
func Error(format string, v ...interface{}) {
    log.Errorf(format, v...)
}

var (

    // ConfigError ...
    ConfigError = "%v type=config.error"

    // HTTPError ...
    HTTPError = "%v type=http.error"

    // HTTPWarn ...
    HTTPWarn = "%v type=http.warn"

    // HTTPInfo ...
    HTTPInfo = "%v type=http.info"
  )
  */
