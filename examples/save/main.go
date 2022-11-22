package main

import (
	"fmt"
	"net/http"
	"time"

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

		thumb.Save(fmt.Sprintf("%d.png", time.Now().Unix()))

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	})

	http.ListenAndServe(":5000", nil)
}
