# go-imageupload

[![GoDoc](https://godoc.org/github.com/olahol/go-imageupload?status.svg)](https://godoc.org/github.com/olahol/go-imageupload)

> :white_square_button: Gracefully handle image uploading and thumbnail creation.

## Install

```bash
go get github.com/olahol/go-imageupload
```

## [Example](https://github.com/olahol/go-imageupload/tree/master/examples)

Thumbnail creator using [Gin](https://github.com/gin-gonic/gin).

```go
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
```

```html
<html>
<body>
  <form method="POST" action="/upload" enctype="multipart/form-data">
    <input type="file" name="file">
    <input type="submit">
  </form>
</body>
</html>
```

## [Documentation](https://godoc.org/github.com/olahol/go-imageupload)
