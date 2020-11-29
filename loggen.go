// loggen
//
// Author:  Ondrej Sika <ondrej@ondrejsika.com>
// Source:  https://github.com/ondrejsika/git-wip
// License: MIT
//

package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	logFile := flag.String("log-file", "", "logfile, default is STDOUT")
	logDebug := flag.Bool("debug", false, "generate DEBUG log messages")
	logInfo := flag.Bool("info", false, "generate INFO log messages")
	logWarn := flag.Bool("warn", false, "generate WARN log messages")
	logError := flag.Bool("error", false, "generate ERROR log messages")

	flag.Parse()

	// Log to logFile or STDOUT
	var logger *log.Logger
	if *logFile != "" {
		f, err := os.OpenFile(*logFile,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		log.Println("Logging into file" + *logFile)
		logger = log.New(f, "loggen ", log.LstdFlags)
		logger.Println("Logging into file" + *logFile)

	} else {
		log.Println("Logging into STDOUT")
		logger = log.New(os.Stdout, "loggen ", log.LstdFlags)
		logger.Println("Logging into STDOUT")
	}

	if !(*logDebug || *logInfo || *logWarn || *logError) {
		logger.Println("ERROR No logging output enabled. See \"loggen -h\"")
		os.Exit(1)
	}

	for {
		time.Sleep(time.Second)

		randomNumber := rand.Intn(100)
		if randomNumber > 90 && *logError {
			logger.Println("ERROR An error is usually an exception that has been caught and not handled.")
			continue
		}
		if randomNumber > 70 && *logWarn {
			logger.Println("WARN A warning that should be ignored is usually at this level and should be actionable.")
			continue
		}
		if randomNumber > 30 && *logInfo {
			logger.Println("INFO This is less important than debug log and is often used to provide context in the current task.")
			continue
		}
		if *logDebug {
			logger.Println("DEBUG This is a debug log that shows a log that can be ignored.")
			continue
		}
	}
}
