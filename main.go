package main

import (
	"fmt"

	"github.com/hcbt/vega/files"
	"github.com/hcbt/vega/analysis"
)

func main(){
	videoID := "nbxPa9Y39JE"

	fmt.Println("Downloading video:", videoID)
	videofilename, _ := files.DownloadVideo(videoID)
	fmt.Println("Converting video to audio:", videoID)
	audiofilename, _ := files.ConvertAudio(videofilename, videoID)
	//files.ConvertAudio(videofilename, videoID)
	//fmt.Println("Rendering spectrogram:", videoID)
	analysis.GenerateSpectrogram(audiofilename)
}
