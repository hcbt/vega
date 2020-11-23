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
	ch1, ch2, _ := analysis.ReadWav(audiofilename)
	fmt.Println("Samples saved:", videoID)

	for value := range ch1 {
		fmt.Printf("%d %d\n", ch1[value], ch2[value])
		//fmt.Printf("%d %d\n", value, ch2[value:])
	}
}
