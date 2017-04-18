/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package log

import (
	"log"
	"os"
	"time"

	"github.com/SiCo-DevOps/cfg"
)

var (
	filename    string
	warningFile = cfg.Config.Log.WARNING
	errorFile   = cfg.Config.Log.ERROR
	fatalFile   = cfg.Config.Log.FATAL
	logFileDir  = cfg.Config.Log.Logpath
)

func LogMsg(s int8, fn string) {
	msg := "[unknown] "
	switch s {
	case 0:
		msg += "Success"
	default:
		msg += "Undefine"

	}
	LogProduce("unknown", time.Now().Format("2006-01-02 15:04:05")+" "+fn+" "+msg)
}

func LogWarnMsg(s int8, fn string) {
	msg := "[warining] "
	switch s {
	case 0:
		msg += "Success"
	default:
		msg += "Undefine"

	}
	LogProduce("warning", time.Now().Format("2006-01-02 15:04:05")+" "+fn+" "+msg)
}

func LogErrMsg(s int8, fn string) {
	msg := "[Error] "
	switch s {
	case 0:
		msg += "Cannot open redis connect"
	case 1:
		msg += "Cannot open mongo connect"
	case 2:
		msg += "Cannot open mysql connect"
	case 3:
		msg += "Cannot open mq connect"
	case 4:
		msg += "Cannot open grpc connect"
	case 10:
		msg += "Redis operation error"
	case 11:
		msg += "Redis unknown error"
	case 12:
		msg += "Redis value parse error"
	default:
		msg = "Undefine"

	}
	LogProduce("error", time.Now().Format("2006-01-02 15:04:05")+" "+fn+" "+msg)
}

// DO NOT USE THIS FUNC
func LogFatalMsg(s int8, fn string) {
	msg := "[fatal] "
	switch s {
	case 0:
		msg += "Cannot open redis connect"
	default:
		msg = "Undefine"

	}
	LogProduce("fatal", time.Now().Format("2006-01-02 15:04:05")+" "+fn+" "+)
	os.Exit(1)
}

func LogProduce(level string, msg string) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("File cannot Write in " + filename)
		}
	}()
	switch level {
	case "warning":
		filename = warningFile
	case "error":
		filename = errorFile
	case "fatal":
		filename = fatalFile
	default:
		filename = "unknown.log"
	}
	fd, err := os.OpenFile(logFileDir+filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
	defer fd.Close()
	if err != nil {
		log.Panicln(filename + " :Open log file Failed")
	}

	logger := log.New(fd, "[SiCo]", log.Lshortfile)
	logger.Println(msg)
}
