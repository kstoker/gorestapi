package logger

import (
	"flag"
	"os"
	"log"
	"go/build"
	"time"
)

var (
  Log *log.Logger
)

func init() {
	// set location of log file
	t := time.Now()
	var logpath = build.Default.GOPATH + "src/gorestapi/log/info_" + t.Format("20060102150405") + ".log"

	flag.Parse()
	var file, err = os.Create(logpath)

	if err != nil {
		panic(err)
	}
	Log = log.New(file, "", log.LstdFlags|log.Llongfile)
//	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
//	Log = log.New(file, "", log.Ldate|log.Ltime|log.Llongfile)
	Log.Println("LogFile: " + logpath)
}
