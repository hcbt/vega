package backend

import (
	"fmt"
	"os"

	"github.com/hcbt/vega/utils"
	"github.com/hcbt/vega/plot"
)

//Process processes an api request
func Process(tempdir string, videoID string) {
	cwd, _ := os.Getwd()

	os.Chdir(tempdir)

	fmt.Println("Downloading video:", videoID)
	utils.DownloadVideo(videoID)

	fmt.Println("Converting video to audio:", videoID)
	utils.ConvertAudio(videoID)

	fmt.Println("Drawing spectrogram:", videoID)
	plot.Spectrogram(cwd, videoID)
}
