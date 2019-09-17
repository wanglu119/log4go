package log4go

import (

)

type ConsoleLogWriterWL struct {
	*ConsoleLogWriter
	outCategory []string
}

func NewConsoleLogWriterWL() *ConsoleLogWriterWL {
	consoleWriter := &ConsoleLogWriterWL{NewConsoleLogWriter(), []string{}}
	
	return consoleWriter
}

func (c *ConsoleLogWriterWL) SetFormat(format string) {
	c.ConsoleLogWriter.SetFormat(format)
}

func (c *ConsoleLogWriterWL) LogWrite(rec *LogRecord) {
	if len(c.outCategory) >0 {
		for _, catecory := range c.outCategory {
			if catecory == rec.Category {
				c.ConsoleLogWriter.LogWrite(rec)
			}
		}
	} else {
		c.ConsoleLogWriter.LogWrite(rec)
	}
}

func (c *ConsoleLogWriterWL) SetOutCategory(category string) {
	c.outCategory = append(c.outCategory, category)
}

func (c *ConsoleLogWriterWL) Close() {
	c.ConsoleLogWriter.Close()
}


// ---------------------------------------
type ConsoleLogWriterEmpty struct {
	ConsoleLogWriter
}

func NewConsoleLogWriterEmpty() *ConsoleLogWriterEmpty {
	consoleWriter := &ConsoleLogWriterEmpty {
	}
		
	return consoleWriter
}

func (c *ConsoleLogWriterEmpty) LogWrite(rec *LogRecord) {
}

func (c *ConsoleLogWriterEmpty) Close() {
}

// ----------------------------------------------------

// Wrapper for (*Logger).AddFilter
func AddConsoleFilter(name string) {
	Global.AddFilter(name, FINE, NewConsoleLogWriterEmpty())
}

func SetConsoleOutCategory(category string) { 
	consoleWriter := Global["stdout"].LogWriter.(*ConsoleLogWriterWL)
	consoleWriter.SetOutCategory(category)
}

func SetConsoleLogLevel(lvl Level) {
	Global["stdout"].Level = lvl
}

// Create a new logger with a "stdout" filter configured to send log messages at
// or above lvl to standard output.
func NewWLDefaultLogger(lvl Level) Logger {
	return Logger{
		"stdout": &Filter{lvl, NewConsoleLogWriterWL(), "DEFAULT"},
	}
}



