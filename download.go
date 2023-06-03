package main

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/schollz/progressbar/v3"

	"github.com/briandowns/spinner"

	"github.com/YakirOren/sdarot"
	log "github.com/sirupsen/logrus"
)

const SpinnerDuration = 100 * time.Millisecond

type Downloader struct {
	client *sdarot.Client
}

func (d *Downloader) Download(ep sdarot.VideoRequest, seriesEnglishName string) {
	dirPath := filepath.Join(seriesEnglishName, strconv.Itoa(ep.Season))

	episodePath := filepath.Join(dirPath, strconv.Itoa(ep.Episode)+".mp4")

	if _, err := os.Stat(episodePath); !errors.Is(err, os.ErrNotExist) {
		log.Println(episodePath, " already exists!")
		return
	}

	err := os.MkdirAll(dirPath, 0777)
	if err != nil {
		log.Fatal("could not create dir")
	}

	for {
		log.Println("Getting video: ", ep.Episode)
		spin := spinner.New(spinner.CharSets[9], SpinnerDuration) // Build our new spinner

		spin.Start()
		video, videoErr := d.client.GetVideo(ep)
		spin.Stop()

		if errors.Is(videoErr, sdarot.ErrServerOverLoad) {
			log.Println("servers are in overload")
			continue
		}

		if videoErr != nil {
			log.Fatal(videoErr)
		}

		done := d.download(episodePath, video)
		if !done {
			continue
		}

		log.Println("\nDone")

		break
	}
}

func (d *Downloader) download(episodePath string, video *sdarot.Video) bool {
	file, err := os.Create(episodePath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Downloading", file.Name())

	bar := progressbar.DefaultBytes(
		video.Size,
		"Downloading...",
	)

	b := io.MultiWriter(file, bar)

	if downloadErr := d.client.Download(video, b); downloadErr != nil {
		log.Error("the download failed unexpectedly\n")
		return false
	}

	return true
}
