package analysis

import (
	"io"
	"os"

	"github.com/youpy/go-wav"
)

//ReadWav - reads .wav file and loads it into array
func ReadWav(filename string) ([]int, []int, error) {
	ch1 := []int{}
	ch2 := []int{}

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
			ch1 = append(ch1, reader.IntValue(sample, 0))
			ch2 = append(ch2, reader.IntValue(sample, 1))
		}
	}
	return ch1, ch2, nil
}