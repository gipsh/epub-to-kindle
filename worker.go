package main

import (

	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

)

var EBOOK_CONVERT_PATH = "/usr/bin/ebook-convert"


type EBook struct {
	Email string
	Epub string
	Mobi string
}

func Get_MOBI_Target(epub string) string {
	name := strings.TrimSuffix(epub, filepath.Ext(epub))
	return fmt.Sprintf("%s.mobi",name)

}

func Convert_Run(ebook EBook, cfg Config) {

	fmt.Println("converting")
	ebook.Mobi = Get_MOBI_Target(ebook.Epub)

	out, err := exec.Command(EBOOK_CONVERT_PATH, ebook.Epub, ebook.Mobi).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The date is %s\n", out)

	Send_Mail(ebook, cfg)
}

func StartWorker(ch <-chan EBook, cfg Config) {

	for {
	  ebook := <-ch
	  Convert_Run(ebook, cfg)
	}

}


func Convert_RunEx(epub string) {


        fmt.Println("converting")
        mobi := Get_MOBI_Target(epub)

        cmd := exec.Command(EBOOK_CONVERT_PATH, epub, mobi)

	stdout, err := cmd.StdoutPipe()

	fmt.Println(stdout)

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("end\n")

}
