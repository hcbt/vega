package utils

import (
	"testing"
)

func TestDownloadVideo(t *testing.T) {
	type args struct {
		videoID string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"54e6lBE3BoQ", args{videoID: "54e6lBE3BoQ"}, "54e6lBE3BoQ.mp4", false},
		{"n3kPvBCYT3E", args{videoID: "n3kPvBCYT3E"}, "n3kPvBCYT3E.mp4", false},
		{"nbxPa9Y39JE", args{videoID: "nbxPa9Y39JE"}, "nbxPa9Y39JE.mp4", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DownloadVideo(tt.args.videoID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadVideo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DownloadVideo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertAudio(t *testing.T) {
	type args struct {
		inputfilename string
		videoID       string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"54e6lBE3BoQ", args{inputfilename: "54e6lBE3BoQ.mp4", videoID: "54e6lBE3BoQ"}, "54e6lBE3BoQ.wav", false},
		{"n3kPvBCYT3E", args{inputfilename: "n3kPvBCYT3E.mp4", videoID: "n3kPvBCYT3E"}, "n3kPvBCYT3E.wav", false},
		{"nbxPa9Y39JE", args{inputfilename: "nbxPa9Y39JE.mp4", videoID: "nbxPa9Y39JE"}, "nbxPa9Y39JE.wav", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertAudio(tt.args.inputfilename, tt.args.videoID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertAudio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertAudio() = %v, want %v", got, tt.want)
			}
		})
	}
}
