package main

import (
	"net/http"

	"github.com/olahol/go-imageupload"
)

var currentImage *imageupload.Image

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
		if currentImage == nil {
			http.NotFound(w, r)
			return
		}

		currentImage.Write(w)
	})

	http.HandleFunc("/thumbnail", func(w http.ResponseWriter, r *http.Request) {
		if currentImage == nil {
			http.NotFound(w, r)
			return
		}

		t, err := imageupload.ThumbnailJPEG(currentImage, 300, 300, 80)

		if err != nil {
			panic(err)
		}

		t.Write(w)
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

		currentImage = img

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	})

	http.ListenAndServe(":5000", nil)
}
