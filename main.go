package main

import (
	"os"
	"fmt"
	"flag"
)

func main() {
        var cfg Config
        readFile(&cfg)

	web := flag.Bool("web", false, "enable web mode")

	flag.Parse()

	if *web {
	  ch := make(chan EBook, 1)
          go StartWorker(ch, cfg)
          StartWeb(ch, cfg)
	} else {
          if len(os.Args) <= 2 {
                fmt.Fprintln(os.Stderr, "You must specify a file and an email address or -web mode")
		flag.Usage()
                os.Exit(1)
          }

          ebook := EBook{ Epub: os.Args[1],
                      Email: os.Args[2],
                      Mobi: "", }

          Convert_Run(ebook, cfg)

        }


}
