package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var configurationFile = flag.String("config", "", "the mbsync configuration file to use")
	var account = flag.String("account", "", "the account in the configuration file to use")
	var cmd = flag.String("command", "", "the command to run when new messages arrive")
	flag.Parse()

	if *configurationFile == "" || *account == "" || *cmd == "" {
		fmt.Println("Required configuration parameters not provided:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	config := parseConfiguration(*configurationFile, *account)
	config.cmd = *cmd

	runIdle(config, "INBOX")
}
