package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type logger struct {
	fileName string
	*log.Logger
}

var once sync.Once
var Logger = getLogger()

func getLogger() *logger {
	var logger *logger
	once.Do(func() {
		logger = setLogger("./log/err.log")
	})
	return logger
}

func setLogger(filePath string) *logger {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)

	if err != nil {
		fmt.Println(err)
	}

	log.SetOutput(file)
	return &logger{
		fileName: file.Name(),
		Logger:   log.New(file, "school-helper: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
