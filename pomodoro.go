// Pomodoro timer
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
	"strings"

	"github.com/deckarep/gosx-notifier"
	"github.com/dustin/go-humanize"
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't open display:", err)
		os.Exit(2)
	}
	defer termbox.Close()

	onBreak := false

	for {
		blockName := "pomodoro"
		notificationTitle := "Pomodoro done.  Time to take a break!"
		start := time.Now()

		wait, err := waitDuration()
		if err != nil {
			flag.Usage()
			os.Exit(2)
		}

		breakTime, err := breakDuration()
		if err != nil {
			flag.Usage()
			os.Exit(2)
		}

		if onBreak {
			wait, breakTime = breakTime, wait
			blockName = "break"
			notificationTitle = "Break done.  Time to continue working!"
		}

		finish := start.Add(wait)

		formatter := formatSeconds
		switch {
		case wait >= 24*time.Hour:
			formatter = formatDays
		case wait >= time.Hour:
			formatter = formatHours
		case wait >= time.Minute:
			formatter = formatMinutes
		}

		if *simple {
			fmt.Printf("Start %s for %s.\n\n", blockName, wait)
			simpleCountdown(finish, formatter)
		} else {
			fullscreenCountdown(start, finish, formatter, strings.ToUpper(blockName))
		}

		if !*hidden {
			note := gosxnotifier.NewNotification(fmt.Sprintf("Your next %s will start %s", blockName, humanize.Time(time.Now().Add(breakTime+1000))))
			note.Title = notificationTitle
			note.AppIcon = "tomato.png"
			err := note.Push()
			if err != nil {
				log.Println("Could not send notification.")
			}
		}

		if !*silent {
			fmt.Println("\a") // \a is the bell literal.
		}

		onBreak = !onBreak
	}
}
