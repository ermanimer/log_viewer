package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ermanimer/color"
)

//prefixes
const (
	debugPrefix   = "Debug"
	infoPrefix    = "Info"
	warningPrefix = "Warning"
	errorPrefix   = "Error"
	fatalPrefix   = "Fatal"
)

//default filename and default prefixes
const (
	filename = "default.log"
	prefixes = "diwef"
)

//new line character for Linux
const (
	newLine = "\n"
)

type Log struct {
	datetime       string
	prefix         string
	callerFunction string
	message        string
}

var longestPrefixLength int
var longestCallerFunctionLength int

func main() {
	filename := flag.String("f", filename, "Filename")
	prefixes := flag.String("p", prefixes, "Prefixes to view.")
	flag.Parse()
	viewLogs(*filename, *prefixes)
}

func viewLogs(filename string, prefixes string) {
	//open log file
	logFile, err := os.Open(filename)
	if err != nil {
		printError("Opening log file failed!")
		return
	}
	defer func() {
		err = logFile.Close()
		if err != nil {
			printError("Closing log file failed!")
		}
	}()
	//read lines from log file
	var logs []Log
	longestPrefixLength = 0
	longestCallerFunctionLength = 0
	reader := bufio.NewReader(logFile)
	for {
		bytesOfLine, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		line := string(bytesOfLine)
		//parse log
		log := parseLog(line)
		if log == nil {
			printError("Parsing log failed!")
			return
		}
		//filter log
		viewLog := false
		if log.prefix == debugPrefix && strings.Contains(prefixes, "d") {
			viewLog = true
		}
		if log.prefix == infoPrefix && strings.Contains(prefixes, "i") {
			viewLog = true
		}
		if log.prefix == warningPrefix && strings.Contains(prefixes, "w") {
			viewLog = true
		}
		if log.prefix == errorPrefix && strings.Contains(prefixes, "e") {
			viewLog = true
		}
		if log.prefix == fatalPrefix && strings.Contains(prefixes, "f") {
			viewLog = true
		}
		if viewLog {
			logs = append(logs, *log)
			updateLongestPrefixLength(log.prefix)
			updateLongestCallerFunctionLength(log.callerFunction)
		}
	}
	//view log
	for _, log := range logs {
		viewLog(&log)
	}
}

func parseLog(line string) *Log {
	line = strings.Trim(line, "[]")
	segments := strings.Split(line, "][")
	if len(segments) != 4 {
		return nil
	}
	log := Log{
		datetime:       segments[0],
		prefix:         segments[1],
		callerFunction: segments[2],
		message:        segments[3],
	}
	return &log
}

func updateLongestPrefixLength(prefix string) {
	prefixLength := len(prefix)
	if prefixLength > longestPrefixLength {
		longestPrefixLength = prefixLength
	}
}

func updateLongestCallerFunctionLength(callerFunction string) {
	callerFunctionLength := len(callerFunction)
	if callerFunctionLength > longestCallerFunctionLength {
		longestCallerFunctionLength = callerFunctionLength
	}
}

func viewLog(log *Log) {
	setPrefixLength(log)
	setPrefixColor(log)
	setCallerFunctionLength(log)
	setCallerFunctionColor(log)
	fmt.Printf("[%v][%v][%v]: %v%v", log.datetime, log.prefix, log.callerFunction, log.message, newLine)
}

func setPrefixLength(log *Log) {
	padding := strings.Repeat(" ", longestPrefixLength-len(log.prefix))
	log.prefix = fmt.Sprintf("%v%v", log.prefix, padding)
}

func setCallerFunctionLength(log *Log) {
	padding := strings.Repeat(" ", longestCallerFunctionLength-len(log.callerFunction))
	log.callerFunction = fmt.Sprintf("%v%v", log.callerFunction, padding)
}

func setPrefixColor(log *Log) {
	switch strings.TrimSpace(log.prefix) {
	case debugPrefix:
		log.prefix = color.Cyan(log.prefix)
	case infoPrefix:
		log.prefix = color.Green(log.prefix)
	case warningPrefix:
		log.prefix = color.Yellow(log.prefix)
	case errorPrefix:
		log.prefix = color.Red(log.prefix)
	case fatalPrefix:
		log.prefix = color.Magenta(log.prefix)
	}
}

func setCallerFunctionColor(log *Log) {
	log.callerFunction = color.Blue(log.callerFunction)
}

func printError(message string) {
	fmt.Printf("log-viewer: %s%s", message, newLine)
}
