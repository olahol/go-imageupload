package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})

	r.POST("/upload", func(c *gin.Context) {
		img, err := imageupload.Process(c.Request, "file")

		if err != nil {
			panic(err)
		}

		thumb, err := imageupload.ThumbnailPNG(img, 300, 300)

		if err != nil {
			panic(err)
		}

		thumb.Write(c.Writer)
	})

	r.Run(":5000")
}
