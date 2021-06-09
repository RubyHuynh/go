package log

import (
	"log"
	"os"
	"sync"
)

type logger struct {
	filename string
	*log.Logger
}

var _log *logger
var _once sync.Once

// start loggeando
func GetInstance() *logger {
	_once.Do(func() {
		path,err := os.Getwd()
		if err != nil {
			path = "./"
		}
		_log = createLogger(path + "/full.log")
	})
	return _log
}

func createLogger(fname string) *logger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	return &logger{
		filename: fname,
		Logger:   log.New(file, "hthngoc ", log.Ldate | log.Ltime),
	}
}



