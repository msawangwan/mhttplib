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
	entry *log.Logger
}

func (al *ProxyLogger) LogStatus(m string) {
	al.entry.Printf("[status : %s]", m)
}

func (al *ProxyLogger) LogRecord(record ProxyRequestRecord) {
	al.entry.Printf(
		"[%s accessed from %s]\n[%s : %s]\n[%s : %s]\n[%d : %d]\n[%f : %f]",
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
		entry: log.New(
			io.MultiWriter(logfile, os.Stdout),
			LOG_MSG_PREFIX,
			log.Ldate|log.Ltime|log.Lshortfile,
		),
	}
}
