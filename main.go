package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var configurationFile = flag.String("config", "", "the mbsync configuration file to use")
	var account = flag.String("account", "", "the account in the configuration file to use")
	flag.Parse()

	if *configurationFile == "" || *account == "" {
		fmt.Println("Required configuration parameters not provided:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("conf: %#v\n", parseConfiguration(*configurationFile, *account))
}
