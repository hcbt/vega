package analysis

import (
	"fmt"
	"io"
	"os"
	"github.com/youpy/go-wav"
)

func GenerateSpectrogram(filename string)(error){
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	reader := wav.NewReader(file)

	defer file.Close()

	for {
		samples, err := reader.ReadSamples()
		if err == io.EOF {
			break
		}
		for _, sample := range samples {
			fmt.Printf("L/R: %d/%d\n", reader.IntValue(sample, 0), reader.IntValue(sample, 1))
		}
	}

	return nil
}