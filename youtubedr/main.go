package main

import (
	"flag"
	"log"
	. "../../youtubedr"
)

func main() {
	flag.Parse()
	currentDir := flag.Arg(1)
	log.Println("download to dir=", currentDir)
	y := NewYoutube(true)
	arg := flag.Arg(0)
	y.DecodeURL(arg)
	y.StartDownload(currentDir)
}
