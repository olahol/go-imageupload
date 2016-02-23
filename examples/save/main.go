package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
	"net/http"
	"time"
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

		thumb.Save(fmt.Sprintf("%d.png", time.Now().Unix()))

		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.Run(":5000")
}
