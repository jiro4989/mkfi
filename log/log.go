package log

import (
	"log"

	"github.com/jiro4989/mkfi/global"
)

func logging(level string, v ...interface{}) {
	var v2 []interface{}
	v2 = append(v2, level)
	v2 = append(v2, v...)
	log.Println(v2)
}

func Debug(v ...interface{}) {
	if global.DebugFlag {
		logging("[DEBUG]", v...)
	}
}

func Info(v ...interface{}) {
	logging("[INFO]", v...)
}

func Error(v ...interface{}) {
	logging("[ERROR]", v...)
}
