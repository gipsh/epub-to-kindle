package main

import (
	"os"
	"fmt"
)

func main() {
        var cfg Config
        readFile(&cfg)

	fmt.Println(cfg.SMTP)

        if len(os.Args) <= 2 {
                fmt.Fprintln(os.Stderr, "You must specify a file and an email address")
                os.Exit(1)
        }

	ch := make(chan EBook, 1)

	go StartWorker(ch, cfg)

	StartWeb(ch, cfg)

//	ebook := EBook{ Epub: os.Args[1],
//			Email: os.Args[2],
//			Mobi: "", }

//	Convert_Run(ebook)
}
