package util

import (
	"log"
	"os"
)

var ErrorLog *log.Logger

func SetLog(fs *os.File) {
	ErrorLog = log.New(fs, "[error]", 5)
}
