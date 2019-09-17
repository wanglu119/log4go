package main 

import (
	
	log "log4go"
)

func main() {
	log.AddConsoleFilter("wl")
	log.SetConsoleOutCategory("wl")
	log.SetConsoleLogLevel(log.DEBUG)
	
	log.Debug("test")
	
	log.LOGGER("wl").Debug("xxxxxx")
	log.LOGGER("wl").Info("yyyyyyy")
	
	log.Close()
}

