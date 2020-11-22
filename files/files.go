package files

import (
	"io"
	"os"
	"log"
	"fmt"
	"github.com/kkdai/youtube/v2"
	ffmpeg "github.com/floostack/transcoder/ffmpeg"
)

//DownloadVideo - downloads video by id
func DownloadVideo(videoID string) (string, error){
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		return "", err
	}

	resp, err := client.GetStream(video, &video.Formats[0])
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	file, err := os.Create(fmt.Sprintf("%s.mp4", videoID))
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}
	
	return file.Name(), nil
}

//ConvertAudio - converts dowloaded youtube video to .wav
func ConvertAudio(inputfilename string, videoID string) (string, error){
	format := "wav"
	overwrite := true
	outputfilename := videoID + ".wav"

	opts := ffmpeg.Options{
		OutputFormat: &format,
		Overwrite: &overwrite,
	}

	ffmpegConf := &ffmpeg.Config{
		FfmpegBinPath: "/usr/local/bin/ffmpeg",
		FfprobeBinPath: "/usr/local/bin/ffprobe",
		ProgressEnabled: true,
	} 
	
	progress, err := ffmpeg.
		New(ffmpegConf).
		Input(inputfilename).
		Output(outputfilename).
		WithOptions(opts).
		Start(opts)

	if err != nil {
		return outputfilename, err
	}

	for msg := range progress {
		log.Printf("%+v", msg)
	}

	err = os.Remove(inputfilename)
	if err != nil {
		return outputfilename, err
	}

	return outputfilename, nil
}