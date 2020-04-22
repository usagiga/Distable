package main

import "flag"

func ParseArgs() (useStream bool) {
	us := flag.Bool("stream", false, "Use stream-mode sync (Sync emoji on update soon)")

	flag.Parse()

	return *us
}
