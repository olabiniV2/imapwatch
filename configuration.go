package main

import (
	"bufio"
	"os"
	"strings"
)

type accountInformation struct {
	configurationFile string
	host              string
	port              string
	username          string
	password          string
	cmd               string
}

func parseConfiguration(file, name string) *accountInformation {
	confs := make(map[string]*accountInformation)
	var ai *accountInformation = nil

	f, _ := os.Open(file)
	defer f.Close()

	readingAccount := false

	s := bufio.NewScanner(f)
	for s.Scan() {
		t := strings.SplitN(strings.TrimSpace(s.Text()), " ", 2)
		if t != nil && len(t) > 1 && t[0][0] != '#' {
			instruction := t[0]
			rest := t[1]

			switch instruction {
			case "IMAPAccount":
				readingAccount = true
				ai = &accountInformation{configurationFile: file}
				confs[rest] = ai
			case "IMAPStore":
				readingAccount = false
			case "MaildirStore":
				readingAccount = false
			case "Channel":
				readingAccount = false
			case "Host":
				if readingAccount {
					ai.host = rest
				}
			case "Port":
				if readingAccount {
					ai.port = rest
				}
			case "User":
				if readingAccount {
					ai.username = rest
				}
			case "Pass":
				if readingAccount {
					ai.password = rest
				}
			default:
			}
		}
	}

	return confs[name]
}
