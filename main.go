package main

import (
	//"fmt"

	"github.com/hcbt/vega/analysis"
	//"github.com/hcbt/vega/utils"
	"github.com/hcbt/vega/stft"
	"github.com/hcbt/vega/plot"
)

func main() {
	//videoID := "kiQtEDrdYDY"
	audiofilename := "sine.wav"

	//fmt.Println("Downloading video:", videoID)
	//videofilename, _ := utils.DownloadVideo(videoID)

	//fmt.Println("Converting video to audio:", videoID)
	//audiofilename, _ := utils.ConvertAudio(videofilename, videoID)

	//fmt.Println("Getting samples:", videoID)
	ch1, _, _ := analysis.ReadWav(audiofilename)

	//fmt.Println("Converting samples to STFT: ", videoID)
	values := stft.GenStft(ch1)

	//fmt.Println("Plotting spectrogram: ", videoID)
	plot.PlotSpectrogram(values)
}
