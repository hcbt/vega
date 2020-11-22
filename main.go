package main

import (
	"fmt"

	"github.com/hcbt/vega/analysis"
	"github.com/hcbt/vega/utils"
)

func main() {
	videoID := "kiQtEDrdYDY"

	fmt.Println("Downloading video:", videoID)
	videofilename, _ := utils.DownloadVideo(videoID)

	fmt.Println("Converting video to audio:", videoID)
	audiofilename, _ := utils.ConvertAudio(videofilename, videoID)

	fmt.Println("Getting samples:", videoID)
	sampledata, _ := analysis.ReadWav(audiofilename)
	fmt.Println("Samples saved:", videoID)

	fmt.Println("Samples to FFT:", videoID)
	analysis.WavToFFT(sampledata)
	fmt.Println("FFT saved:", videoID)

	//fmt.Println(sampledata)
}
