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
	logPrefix := flag.String("log-prefix", "loggen", "prefix for logger, default loggen")
	sleepTime := flag.Int("sleep-time", 1000, "prefix for logger, default loggen")
	noDebug := flag.Bool("no-debug", false, "don't generate DEBUG log messages")
	noInfo := flag.Bool("no-info", false, "don't generate INFO log messages")
	noWarn := flag.Bool("no-warn", false, "don't generate WARN log messages")
	noError := flag.Bool("no-error", false, "don't generate ERROR log messages")

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
		logger = log.New(f, *logPrefix+" ", log.LstdFlags)
		logger.Println("Logging into file " + *logFile)

	} else {
		log.Println("Logging into STDOUT")
		logger = log.New(os.Stdout, *logPrefix+" ", log.LstdFlags)
		logger.Println("Logging into STDOUT")
	}

	if *noError && *noWarn && *noInfo && *noDebug {
		logger.Println("ERROR No logging output enabled. See \"loggen -h\"")
		os.Exit(1)
	}

	for {
		time.Sleep(time.Duration(*sleepTime) * time.Millisecond)

		randomNumber := rand.Intn(100)
		if randomNumber > 90 && !*noError {
			logger.Println("ERROR An error is usually an exception that has been caught and not handled.")
			continue
		}
		if randomNumber > 70 && !*noWarn {
			logger.Println("WARN A warning that should be ignored is usually at this level and should be actionable.")
			continue
		}
		if randomNumber > 30 && !*noInfo {
			logger.Println("INFO This is less important than debug log and is often used to provide context in the current task.")
			continue
		}
		if !*noDebug {
			logger.Println("DEBUG This is a debug log that shows a log that can be ignored.")
			continue
		}
	}
}
