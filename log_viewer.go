package main

import (
	"bufio"
	"errors"
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
	datetime string
	prefix   string
	message  string
}

var longestPrefixLength int

func main() {
	filename := flag.String("f", filename, "filename")
	prefixes := flag.String("p", prefixes, "prefixes to view")
	flag.Parse()
	viewLogs(*filename, *prefixes)
}

func viewLogs(filename string, prefixes string) {
	//open log file
	logFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("opening log file failed")
		return
	}
	defer func() {
		err = logFile.Close()
		if err != nil {
			fmt.Println("closing log file failed")
		}
	}()
	//read lines from log file
	var logs []Log
	longestPrefixLength = 0
	reader := bufio.NewReader(logFile)
	for {
		bytesOfLine, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		line := string(bytesOfLine)
		//parse log
		log, err := parseLog(line)
		if err != nil {
			fmt.Println(err.Error())
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
		}
	}
	//view log
	for _, log := range logs {
		viewLog(&log)
	}
}

func parseLog(line string) (*Log, error) {
	segments := strings.Split(line[1:len(line)-1], "][")
	if len(segments) != 3 {
		return nil, errors.New("parsing log failed")
	}
	log := Log{
		datetime: segments[0],
		prefix:   segments[1],
		message:  segments[2],
	}
	return &log, nil
}

func updateLongestPrefixLength(prefix string) {
	prefixLength := len(prefix)
	if prefixLength > longestPrefixLength {
		longestPrefixLength = prefixLength
	}
}

func viewLog(log *Log) {
	setPrefixLength(log)
	setPrefixColor(log)
	fmt.Printf("[%s][%s]%s%s", log.datetime, log.prefix, log.message, newLine)
}

func setPrefixLength(log *Log) {
	padding := strings.Repeat(" ", longestPrefixLength-len(log.prefix))
	log.prefix = fmt.Sprintf("%s%s", log.prefix, padding)
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
		log.prefix = color.Red(log.prefix)
	}
}
