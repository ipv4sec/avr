package logger

import (
	"fmt"
	"os"
	"time"
)

func Info(msg ...interface{}) {
	_, err := fmt.Fprintln(os.Stdout,time.Now().String()[:19], msg)
	if err != nil {
		panic(err)
	}
}