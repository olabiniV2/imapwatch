package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	idle "github.com/emersion/go-imap-idle"
	"github.com/emersion/go-imap/client"
)

func idleOnClient(c *client.Client, tup chan<- bool) {
	updates := make(chan client.Update)
	c.Updates = updates

	done := make(chan error, 1)

	go func() {
		ic := idle.NewClient(c)
		ic.LogoutTimeout = 5 * time.Minute
		done <- ic.IdleWithFallback(nil, 0)
	}()

	for {
		select {
		case <-updates:
			tup <- true
		case err := <-done:
			if err != nil {
				log.Println(err)
			}
			return
		}
	}
}

func runUpdate(ac *accountInformation, box string) {
	cmd := exec.Command(ac.cmd, box)
	cmd.Stdout = os.Stdout
	e := cmd.Run()
	if e != nil {
		fmt.Printf("  error when running command: %v\n", e)
	}
}

func runIdle(ac *accountInformation, box string) {
	timeForUpdate := make(chan bool, 1000)
	go func() {
		t := false
		for {
			select {
			case <-timeForUpdate:
				t = true
			case <-time.After(10 * time.Second):
				if t {
					t = false
					runUpdate(ac, box)
				}
			}
		}
	}()

	c, err := client.DialTLS(fmt.Sprintf("%s:%s", ac.host, ac.port), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Logout()

	if err := c.Login(ac.username, ac.password); err != nil {
		log.Fatal(err)
	}

	if _, err := c.Select(box, false); err != nil {
		log.Fatal(err)
	}

	for {
		idleOnClient(c, timeForUpdate)
	}
}
