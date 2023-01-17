package corelibs

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func CommonLogger() *logrus.Logger {
	log := logrus.New()
	// path := "./logs/common.log"
	/* Log rotation related functions
	`WithLinkName` establishes a soft link for the latest log
	`WithRotationTime` sets the time for log splitting, how often
	Only one of WithMaxAge and WithRotationCount can be set
	  `WithMaxAge` sets the maximum save time before file cleanup
	  `WithRotationCount` sets the maximum number of saved files before cleaning
	*/
	// The following configuration log rotates a new file every 1 day, keeps the log files of the last 30 day, and automatically clears the excess.
	// writer, _ := rotatelogs.New(
	// 	path+".%Y%m%d",
	// 	rotatelogs.WithLinkName(path),
	// 	rotatelogs.WithMaxAge(time.Duration(30*24*3600)*time.Second),
	// 	rotatelogs.WithRotationTime(time.Duration(24*3600)*time.Second),
	// )

	log.SetReportCaller(true)
	//print to multiple medium
	// log.SetOutput(io.MultiWriter(writer, os.Stdout))
	log.SetOutput(io.MultiWriter(os.Stdout))
	return log
}
