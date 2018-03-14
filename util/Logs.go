package util

import (
	"log"
	"os"
)

var ErrorLog *log.Logger

func SetLogFile(fs *os.File) {
	ErrorLog = log.New(fs, "[error]", 5)
}
