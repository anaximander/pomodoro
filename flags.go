package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

const defaultPomodoroDuration = 15 * time.Minute
const defaultBreakDuration = 5 * time.Minute

var silent = flag.Bool("silent", false, "Don't ring bell after countdown")

var hidden = flag.Bool("hidden", false, "Don't show notification after countdown")

var simple = flag.Bool("simple", false, "Display simple countdown")

func init() {
	const usage = `Usage of pomodoro:

    pomodoro [options] [pomodoroDuration] [breakDuration]

Pomodoro duration defaults to %d minutes.
Break duration defaults to %d minutes.
Durations may be expressed as integer minutes
(e.g. "15") or time with units (e.g. "1m30s" or "90s").

Chimes system bell at the end of the timer, unless -silent is set.
Creates a system notification at the end of the timer, unless -hidden is set.
`
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, int(defaultPomodoroDuration/time.Minute), int(defaultBreakDuration/time.Minute))
		flag.PrintDefaults()
	}
	flag.Parse()
}

func waitDuration() (time.Duration, error) {
	arg := flag.Arg(0)

	if arg == "" {
		return defaultPomodoroDuration, nil
	}

	if n, err := strconv.Atoi(arg); err == nil {
		return time.Duration(n) * time.Minute, nil
	}

	if d, err := time.ParseDuration(arg); err == nil {
		return d, nil
	}

	return 0, errors.New("Couldn't parse pomodoro duration.")
}

func breakDuration() (time.Duration, error) {
	arg := flag.Arg(1)

	if arg == "" {
		return defaultBreakDuration, nil
	}

	if n, err := strconv.Atoi(arg); err == nil {
		return time.Duration(n) * time.Minute, nil
	}

	if d, err := time.ParseDuration(arg); err == nil {
		return d, nil
	}

	return 0, errors.New("Couldn't parse break duration.")
}
