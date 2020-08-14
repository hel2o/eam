package g

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	logFileName = flag.String("log", "var/app.log", "Log file name")
)

func AppLog() {
	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "APP start Failed")
		os.Exit(1)
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}
