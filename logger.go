package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	logDefault = 0
	logInfo    = 1
	logDebug   = 2
)

var (
	loglevel       int
	logfile        *os.File
	logfilename    string
	initDone       bool
	bufferedWriter *bufio.Writer
	err            error
)

func setLogLevel(level int) {
	if initDone {
		loglevel = level
	}
}

func getLogLevel() int {
	return loglevel
}

func writeToLog(log string, level int) {
	if initDone {
		if level <= getLogLevel() {
			bytesAvailable := bufferedWriter.Available()
			if bytesAvailable > len(log) {
				_, err := bufferedWriter.WriteString(log + "\n")
				if err != nil {
					fmt.Println("Fatal : Unable to write to log file")
				}
			} else {
				bufferedWriter.Flush()
				writeToLog(log, level)
			}
		}
	}
}

func initLogging(fname string, level int) {
	if fname != "" {
		logfilename = fname
		logfile, err = os.OpenFile("./"+logfilename, os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("Fatal : Unable init logging. Logs in console...")
			return
		}
		setLogLevel(level)
		bufferedWriter = bufio.NewWriter(logfile)
		initDone = true
	}
}
