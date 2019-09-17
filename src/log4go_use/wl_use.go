package main 

import (
	
	log "log4go"
)

func main() {
	log.LoadConfiguration("./config/log4go_config.json", "json")
	log.AddConsoleFilter("wl")
	
	log.SetConsoleOutCategory("wl")
	log.SetConsoleLogLevel(log.DEBUG)
	
	log.Debug("test")
	
	log.LOGGER("wl").Debug("xxxxxx")
	log.LOGGER("wl").Info("yyyyyyy")
	
	log.LOGGER("Test").Debug("x==================")
	log.LOGGER("TestInfo").Debug("y=================")
	log.LOGGER("TestInfo").Info("z=======================")
	
	log.Close()
}

