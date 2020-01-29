package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var UPLOAD_DIR = "upload"

func StartWeb(ch chan<- EBook, cfg Config) {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/", "./public")
	router.POST("/upload", func(c *gin.Context) {
		email := c.PostForm("email")

		// Source
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		filename := fmt.Sprintf("%s/%s", UPLOAD_DIR, filepath.Base(file.Filename))

		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully. I will send it to %s.", filename, email))

		ebook := EBook{ Epub: filename,
	                        Email: email,
                                Mobi: "", }

		ch <- ebook

	})
	router.Run(fmt.Sprintf(":%s",cfg.WEB.Port))
}
