package analysis

import (
	"github.com/mjibson/go-dsp/fft"
	"github.com/youpy/go-wav"
	"io"
	"os"
)

func ReadWav(filename string) ([]float64, error) {
	var sampledata []float64

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	reader := wav.NewReader(file)

	defer file.Close()

	for {
		samples, err := reader.ReadSamples()
		if err == io.EOF {
			break
		}
		for _, sample := range samples {
			sampledata = append(sampledata, reader.FloatValue(sample, 0))
			//fmt.Printf("L/R: %d/%d\n", reader.IntValue(sample, 0), reader.IntValue(sample, 1))
		}
	}

	return sampledata, nil
}

//WavToFFT - returns fourier transform of .wav samples array
func WavToFFT(sampledata []float64) []complex128 {
	fftdata := fft.FFTReal(sampledata)

	return fftdata
}
