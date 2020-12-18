package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	c "github.com/ermanimer/color/v2"
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

type log struct {
	datetime string
	prefix   string
	message  string
}

//color print functions
var (
	printlnError        = (&c.Color{Foreground: 1}).PrintlnFunction()
	printfDatetime      = (&c.Color{Foreground: 240}).PrintfFunction()
	printfDebugPrefix   = (&c.Color{Foreground: 246}).PrintfFunction()
	printfInfoPrefix    = (&c.Color{Foreground: 12}).PrintfFunction()
	printfWarningPrefix = (&c.Color{Foreground: 11}).PrintfFunction()
	printfErrorPrefix   = (&c.Color{Foreground: 9}).PrintfFunction()
	printfFatalPrefix   = (&c.Color{Foreground: 13}).PrintfFunction()
	printlnMessage      = (&c.Color{Foreground: 252}).PrintlnFunction()
)

var prefixLength int

func main() {
	//parse flags
	filename := flag.String("f", filename, "filename")
	prefixes := flag.String("p", prefixes, "prefixes to view")
	flag.Parse()
	//view logs
	viewLogs(*filename, *prefixes)
}

func viewLogs(filename string, prefixes string) {
	//open log file
	logFile, err := os.Open(filename)
	if err != nil {
		printlnError("opening log file failed")
		return
	}
	defer func() {
		err = logFile.Close()
		if err != nil {
			printlnError("closing log file failed")
		}
	}()
	//read lines from log file
	var ls []*log
	r := bufio.NewReader(logFile)
	lineNumber := 0
	for {
		lineNumber++
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		//parse log
		l, err := parseLog(string(line))
		if err != nil {
			printlnError(err.Error(), " on line  ", lineNumber)
			continue
		}
		//filter log
		switch {
		case l.prefix == debugPrefix && strings.Contains(prefixes, "d"):
			ls = append(ls, l)
			updatePrefixLength(l.prefix)
		case l.prefix == infoPrefix && strings.Contains(prefixes, "i"):
			ls = append(ls, l)
			updatePrefixLength(l.prefix)
		case l.prefix == warningPrefix && strings.Contains(prefixes, "w"):
			ls = append(ls, l)
			updatePrefixLength(l.prefix)
		case l.prefix == errorPrefix && strings.Contains(prefixes, "e"):
			ls = append(ls, l)
			updatePrefixLength(l.prefix)
		case l.prefix == fatalPrefix && strings.Contains(prefixes, "f"):
			ls = append(ls, l)
			updatePrefixLength(l.prefix)
		default:
			printlnError("filtering prefix failed on line ", lineNumber)
		}
	}
	//view log
	for _, l := range ls {
		viewLog(l)
	}
}

func parseLog(line string) (*log, error) {
	segments := strings.Split(line[1:len(line)-1], "][")
	if len(segments) != 3 {
		return nil, errors.New("parsing log failed")
	}
	l := &log{
		datetime: segments[0],
		prefix:   segments[1],
		message:  segments[2],
	}
	return l, nil
}

func viewLog(l *log) {
	setPrefixLength(l)
	printfDatetime("%s ", l.datetime)
	switch strings.TrimSpace(l.prefix) {
	case debugPrefix:
		printfDebugPrefix("%s ", l.prefix)
	case infoPrefix:
		printfInfoPrefix("%s ", l.prefix)
	case warningPrefix:
		printfWarningPrefix("%s ", l.prefix)
	case errorPrefix:
		printfErrorPrefix("%s ", l.prefix)
	case fatalPrefix:
		printfFatalPrefix("%s ", l.prefix)
	}
	printlnMessage(l.message)
}

func updatePrefixLength(prefix string) {
	length := len(prefix)
	if length > prefixLength {
		prefixLength = length
	}
}

func setPrefixLength(l *log) {
	padding := strings.Repeat(" ", prefixLength-len(l.prefix))
	l.prefix = fmt.Sprintf("%s%s", l.prefix, padding)
}
