package analysis

import (
	"io"
	"os"
	
	"github.com/youpy/go-wav"
)

//ReadWav - reads .wav file and loads it into array
func ReadWav(filename string) ([]float64, []float64, error) {
	ch1 := []float64{}
	ch2 := []float64{}

	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}

	reader := wav.NewReader(file)

	defer file.Close()

	for {
		samples, err := reader.ReadSamples()
		if err == io.EOF {
			break
		} 
		for _, sample := range samples {
			ch1 = append(ch1, float64(reader.IntValue(sample, 0)))
			ch2 = append(ch2, float64(reader.IntValue(sample, 1)))
		}
	}
	
	return ch1, ch2, nil
}
