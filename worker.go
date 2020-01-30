package main

import (

	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
	"runtime"

)

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
	fmt.Println(ebook)
	ebook.Mobi = Get_MOBI_Target(ebook.Epub)

	ebook_convert := GetConvertPath(cfg)

	out, err := exec.Command(ebook_convert, ebook.Epub, ebook.Mobi).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The date is %s\n", out)

	Send_Mail(ebook, cfg)
}

func GetConvertPath(cfg Config) string {

  switch runtime.GOOS {
    case "linux":
      return cfg.CONVERT.Linux 
    case "darwin":
      return cfg.CONVERT.Darwin
    case "windows": 
      return cfg.CONVERT.Windows
    default:
      return cfg.CONVERT.Linux
    }

}


func StartWorker(ch <-chan EBook, cfg Config) {

	for {
	  ebook := <-ch
	  Convert_Run(ebook, cfg)
	}

}


func Convert_RunEx(epub string, cfg Config) {


        fmt.Println("converting")
        mobi := Get_MOBI_Target(epub)

	ebook_convert := GetConvertPath(cfg)

        cmd := exec.Command(ebook_convert, epub, mobi)

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
