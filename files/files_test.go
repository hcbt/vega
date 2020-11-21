package files

import (
	"testing"
)

func TestDownloadVideo(t *testing.T) {
	var tests = []struct {
		id string
	}{
		{"54e6lBE3BoQ"},
		{"n3kPvBCYT3E"},
		{"nbxPa9Y39JE"},
	}

	for _, tt := range tests {

		file, err := DownloadVideo(tt.id)

		if file == "" {
			t.Errorf("No file created on %v", tt.id)
		}

		if err != nil {
			t.Errorf("got %v on %v", err, tt.id)
		}
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
		wantErr bool
	}{
		{"54e6lBE3BoQ", args{inputfilename: "54e6lBE3BoQ.mp4", videoID: "54e6lBE3BoQ"}, false},
		{"n3kPvBCYT3E", args{inputfilename: "n3kPvBCYT3E.mp4", videoID: "n3kPvBCYT3E"}, false},
		{"nbxPa9Y39JE", args{inputfilename: "nbxPa9Y39JE.mp4", videoID: "nbxPa9Y39JE"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ConvertAudio(tt.args.inputfilename, tt.args.videoID); (err != nil) != tt.wantErr {
				t.Errorf("ConvertAudio() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}