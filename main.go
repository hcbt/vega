package main

import (
	"github.com/hcbt/vega/files"
)

func main() {
	videoID := "nbxPa9Y39JE"

	videofilename, _ := files.DownloadVideo(videoID)
	files.ConvertAudio(videofilename, videoID)
}
