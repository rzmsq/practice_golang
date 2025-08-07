package main

import (
	"log"
	"log/syslog"
)

func main() {
	sysLog, err := syslog.New(syslog.LOG_SYSLOG, "systemLog.go")
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(sysLog)
	log.Print("Evrerything is working fine!")
}
