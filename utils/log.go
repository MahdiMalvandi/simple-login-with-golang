package utils

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	loggerChannel chan string
}

func NewLogger() *Logger {
	l := &Logger{
		loggerChannel: make(chan string, 100),
	}

	go l.start()
	return l
}

func (l *Logger) start() {
	file, err := os.OpenFile("server.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for msg := range l.loggerChannel {
		fullLog := fmt.Sprintf("%s - %s\n", time.Now().Format("2006-01-02 15:04:05"), msg)
		file.WriteString(fullLog)
	}

}

func (l *Logger) Log(message string) {
	l.loggerChannel <- message
}
