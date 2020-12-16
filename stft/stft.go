package stft

import (
	"github.com/hcbt/gossp"
	"github.com/hcbt/gossp/stft"
	"github.com/hcbt/gossp/window"
)

// GenStft - 
func GenStft(data []float64) ([][]float64){
	sr := 44100

	s := &stft.STFT{
		FrameShift: int(float64(sr) / 100.0), // 0.01 sec,
		FrameLen:   2048,
		Window:     window.CreateHanning(2048),
	}

	spectrogram, _ := gossp.SplitSpectrogram(s.STFT(data))

	//rows, cols, values := len(spectrogram), len(spectrogram[0]), spectrogram

	return spectrogram
}