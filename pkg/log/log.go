package log

import (
	"fmt"
	"log"
	"os"
)

func Debugf(format string, v ...any) {
	_ = log.Output(2, fmt.Sprintf("DEBUG "+format+"\n", v...))
}

func Infof(format string, v ...any) {
	_ = log.Output(2, fmt.Sprintf("INFO  "+format+"\n", v...))
}

func Warnf(format string, v ...any) {
	_ = log.Output(2, fmt.Sprintf("WARN  "+format+"\n", v...))
}

func Errorf(format string, v ...any) {
	_ = log.Output(2, fmt.Sprintf("ERROR "+format+"\n", v...))
}

func Fatal(v ...any) {
	_ = log.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}
