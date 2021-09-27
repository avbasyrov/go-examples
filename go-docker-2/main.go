package main

import (
	"github.com/fatih/color"
	"time"
)

func main() {
	for t := range time.Tick(time.Second) {
		color.Magenta("%s\n", t)
	}
}
