// Gracefully handle image uploading and thumbnail creation.
package imageupload

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	_ "image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Image struct {
	Filename    string
	ContentType string
	Data        []byte
	Size        int
}

// Save image to file.
func (i *Image) Save(filename string) error {
	return ioutil.WriteFile(filename, i.Data, 0600)
}

// Convert image to base64 data uri.
func (i *Image) DataURI() string {
	return fmt.Sprintf("data:%s;base64,%s", i.ContentType, base64.StdEncoding.EncodeToString(i.Data))
}

// Write image to HTTP response.
func (i *Image) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", i.ContentType)
	w.Header().Set("Content-Length", strconv.Itoa(i.Size))
	w.Write(i.Data)
}

// Limit the size of uploaded files, put this before imageupload.Process.
func LimitFileSize(maxSize int64, w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxSize)
}

func okContentType(contentType string) bool {
	return contentType == "image/png" || contentType == "image/jpeg" || contentType == "image/gif"
}

// Process uploaded file into an image.
func Process(r *http.Request, field string) (*Image, error) {
	file, info, err := r.FormFile(field)

	if err != nil {
		return nil, err
	}

	contentType := info.Header.Get("Content-Type")

	if !okContentType(contentType) {
		return nil, errors.New(fmt.Sprintf("Wrong content type: %s", contentType))
	}

	bs, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	_, _, err = image.Decode(bytes.NewReader(bs))

	if err != nil {
		return nil, err
	}

	i := &Image{
		Filename:    info.Filename,
		ContentType: contentType,
		Data:        bs,
		Size:        len(bs),
	}

	return i, nil
}

// Create JPEG thumbnail.
func ThumbnailJPEG(i *Image, width int, height int, quality int) (*Image, error) {
	img, _, err := image.Decode(bytes.NewReader(i.Data))

	thumbnail := resize.Thumbnail(uint(width), uint(height), img, resize.Lanczos3)

	data := new(bytes.Buffer)
	err = jpeg.Encode(data, thumbnail, &jpeg.Options{
		Quality: quality,
	})

	if err != nil {
		return nil, err
	}

	bs := data.Bytes()

	t := &Image{
		Filename:    "thumbnail.jpg",
		ContentType: "image/jpeg",
		Data:        bs,
		Size:        len(bs),
	}

	return t, nil
}

// Create PNG thumbnail.
func ThumbnailPNG(i *Image, width int, height int) (*Image, error) {
	img, _, err := image.Decode(bytes.NewReader(i.Data))

	thumbnail := resize.Thumbnail(uint(width), uint(height), img, resize.Lanczos3)

	data := new(bytes.Buffer)
	err = png.Encode(data, thumbnail)

	if err != nil {
		return nil, err
	}

	bs := data.Bytes()

	t := &Image{
		Filename:    "thumbnail.jpg",
		ContentType: "image/png",
		Data:        bs,
		Size:        len(bs),
	}

	return t, nil
}
