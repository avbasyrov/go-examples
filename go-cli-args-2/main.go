package main // go-cli-args

import (
	"flag"
	"fmt"
)

func main() {
	var sourceFile, dstFile string

	flag.StringVar(&sourceFile, "source", "", "source filename")
	flag.StringVar(&sourceFile, "s", "", "source filename (short)")
	flag.StringVar(&dstFile, "destination", "", "destination filename")
	flag.StringVar(&dstFile, "d", "", "destination filename (short)")
	flag.Parse()

	fmt.Println("Source file:", sourceFile)
	fmt.Println("Destination file:", dstFile)
}
