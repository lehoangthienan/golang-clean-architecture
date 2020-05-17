package image

import (
	"bytes"
	"context"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/disintegration/imaging"

	requestModel "github.com/lehoangthienan/marvel-heroes-backend/model/request/image"
	responseModel "github.com/lehoangthienan/marvel-heroes-backend/model/response/image"
	"github.com/lehoangthienan/marvel-heroes-backend/util/errors"
)

type imageService struct{}

// NewService func
func NewService() Service {
	return &imageService{}
}

func (s *imageService) Create(ctx context.Context, req *requestModel.Images) (*responseModel.Images, error) {
	var wg sync.WaitGroup
	var err error
	var rs []string
	ch := make(chan string)

	for i := 0; i < len(req.Images); i++ {
		wg.Add(1)
		go downloadFile("uploads", req.Images[i], &wg, ch)
	}

	for i := 0; i < len(req.Images); i++ {
		rs = append(rs, "/"+<-ch)
	}

	wg.Wait()

	if err != nil {
		return nil, errors.CreateHeroFailedError
	}

	return &responseModel.Images{Images: rs}, err
}

func downloadFile(directory string, url string, wg *sync.WaitGroup, chnl chan string) error {
	defer wg.Done()
	var dst image.Image

	// Get the data
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// decode input data to image
	src, _, err := image.Decode(resp.Body)

	// resize input image
	dst = imaging.Resize(src, 512, 512, imaging.Lanczos)

	//file name
	filepath := directory + "/" + strconv.Itoa(int(time.Now().UnixNano())) + ".jpg"

	chnl <- filepath

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// encode output image to jpeg buffer
	encoded, err := encodeImageToJpg(&dst)

	// Write the body to file
	_, err = io.Copy(out, encoded)

	return err
}

// encode image to jpeg
func encodeImageToJpg(img *image.Image) (*bytes.Buffer, error) {
	encoded := &bytes.Buffer{}
	err := jpeg.Encode(encoded, *img, nil)
	return encoded, err
}

func (s *imageService) GetImageFile(ctx context.Context, path string) (*os.File, error) {
	file, err := os.Open("uploads/" + path)

	if err != nil {
		return nil, err
	}
	return file, nil
}
