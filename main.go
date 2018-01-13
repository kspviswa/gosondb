package main

import (
	"time"
)

func main() {
	initLogging("gsondb.log", logDefault)
	t := time.Now()
	writeToLog(t.Format("Mon Jan _2 15:04:05 2006"), logDefault)
}
