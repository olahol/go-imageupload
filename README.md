# go-imageupload

[![GoDoc](https://godoc.org/github.com/olahol/go-imageupload?status.svg)](https://godoc.org/github.com/olahol/go-imageupload)

> :white_square_button: Gracefully handle image uploading and thumbnail creation.

## Install

```bash
go get github.com/olahol/go-imageupload
```

## [Example](https://github.com/olahol/go-imageupload/tree/master/examples)

Thumbnail creator.

```go
package main

import (
	"net/http"

	"github.com/olahol/go-imageupload"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.NotFound(w, r)
			return
		}

		img, err := imageupload.Process(r, "file")

		if err != nil {
			panic(err)
		}

		thumb, err := imageupload.ThumbnailPNG(img, 300, 300)

		if err != nil {
			panic(err)
		}

		thumb.Write(w)
	})

	http.ListenAndServe(":5000", nil)
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

## Contributors

* Ola Holmström (@olahol)
* Shintaro Kaneko (@kaneshin)


## [Documentation](https://godoc.org/github.com/olahol/go-imageupload)
