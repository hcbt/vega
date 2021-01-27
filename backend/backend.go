package backend

import (
	"fmt"

	"github.com/hcbt/vega/utils"
	"github.com/hcbt/vega/plot"
)

//Process processes an api request
func Process(videoID string) {
	fmt.Println("Downloading video:", videoID)
	videofilename, err := utils.DownloadVideo(videoID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Converting video to audio:", videoID)
	audiofilename, err := utils.ConvertAudio(videofilename, videoID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Drawing spectrogram:", videoID)
	err = plot.Spectrogram(audiofilename, videoID)
	if err != nil {
		fmt.Println(err)
	}
}