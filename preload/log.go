package preload

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func InitLog() {
	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
	// add the calling method as a field
	log.SetReportCaller(true)
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
	// You could set this to any `io.Writer` such as a file
	//logFile, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err != nil {
	//	panic(err)
	//}
	//log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	log.SetOutput(io.MultiWriter(os.Stdout))
}
