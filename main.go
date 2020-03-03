package main

import (
	"bufio"
	"flag"
	"io"
	"os"
	"strings"

	"github.com/ermanimer/color"
)

//prefixes
const (
	debugPrefix   = "Debug  "
	infoPrefix    = "Info   "
	warningPrefix = "Warning"
	errorPrefix   = "Error  "
	fatalPrefix   = "Fatal  "
)

//default filename and prefixes
const (
	defaultFilename = "logs"
	defaultPrefixes = "diwef"
)

func main() {
	filename := flag.String("f", defaultFilename, "Filename")
	prefixes := flag.String("p", defaultPrefixes, "Prefixes to view.")
	flag.Parse()

	viewLogs(*filename, *prefixes)
}

func viewLogs(filename string, prefixes string) {
	//parse prefixes
	viewDebugMessages := strings.Contains(prefixes, "d")
	viewInfoMessages := strings.Contains(prefixes, "i")
	viewWarningMessages := strings.Contains(prefixes, "w")
	viewErrorMessages := strings.Contains(prefixes, "e")
	viewFatalMessages := strings.Contains(prefixes, "f")

	//open log file
	logFile, err := os.Open(filename)
	if err != nil {
		color.Red("Error: Opening log file failed!")
		return
	}
	defer func() {
		err = logFile.Close()
		if err != nil {
			color.Red("Error: Closing log file failed!")
		}
	}()

	//read from log file
	reader := bufio.NewReader(logFile)
	for {
		bytesOfLine, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		line := string(bytesOfLine)

		//view message if prefix is selected
		switch {
		case strings.Contains(line, debugPrefix):
			if viewDebugMessages {
				color.Cyan(line)
			}
		case strings.Contains(line, infoPrefix):
			if viewInfoMessages {
				color.Green(line)
			}
		case strings.Contains(line, warningPrefix):
			if viewWarningMessages {
				color.Yellow(line)
			}
		case strings.Contains(line, errorPrefix):
			if viewErrorMessages {
				color.Red(line)
			}
		case strings.Contains(line, fatalPrefix):
			if viewFatalMessages {
				color.Magenta(line)
			}
		default:
			color.Red("Error: Unknown prefix!")
		}
	}
}
