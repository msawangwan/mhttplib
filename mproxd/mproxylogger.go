package mproxd

import (
	"io"
	"log"
	"os"
)

const (
	LOGFILE_NAME   = "access.log"
	LOG_MSG_PREFIX = "[Proxy Logger] "
)

type ProxyLogger struct {
	logWriter *log.Logger
	logPutter *log.Logger // <-- for printing to console only
}

func (pl *ProxyLogger) LogStatus(m string) {
	pl.logPutter.Printf("\n\t[status : %s]\n", m)
}

func (pl *ProxyLogger) LogRecord(record ProxyRequestRecord) {
	pl.logWriter.Printf(
		"\n\t[%s accessed from %s]\n\t[%s : %s]\n\t[%s : %s]\n\t[%d : %d]\n\t[%f : %f]\n",
		record.RequestedResource,
		record.IP,
		record.CountryName,
		record.CountryCode,
		record.City,
		record.ZipCode,
		record.MetroCode,
		record.AreaCode,
		record.Latitude,
		record.Longitude,
	)
}

func NewProxyLogger() *ProxyLogger {
	logfile, err := os.OpenFile(
		LOGFILE_NAME,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)

	if err != nil {
		log.Printf(
			"error opening %s, log file: %s",
			LOGFILE_NAME,
			err,
		)
	}

	//defer logfile.Close() <- ????

	return &ProxyLogger{
		logWriter: log.New(
			io.MultiWriter(logfile, os.Stdout),
			LOG_MSG_PREFIX,
			log.Ldate|log.Ltime,
		),
		logPutter: log.New(
			os.Stdout,
			LOG_MSG_PREFIX,
			log.Ldate|log.Ltime,
		),
	}
}
