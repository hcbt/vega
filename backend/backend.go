package backend

import (
	"fmt"

	"github.com/hcbt/vega/utils"
	"github.com/hcbt/vega/plot"
)

//Process processes an api request
func Process(videoID string) {
	fmt.Println("Downloading video:", videoID)
	videofilename := utils.DownloadVideo(videoID)

	fmt.Println("Converting video to audio:", videoID)
	audiofilename := utils.ConvertAudio(videofilename, videoID)

	fmt.Println("Drawing spectrogram:", videoID)
	plot.Spectrogram(audiofilename, videoID)
}
