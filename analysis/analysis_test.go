package analysis

import (
	"os"
	"io/ioutil"
	//"reflect"
	"testing"

	"github.com/youpy/go-wav"
)

func TestReadWav (t *testing.T) {
	file, err := os.Open("../tests/a.wav")
	if err != nil {
		t.Fatalf("Failed to load test file")
	}

	reader := wav.NewReader(file)
	fmt, err := reader.Format()
	if err != nil {
		t.Fatal(err)
	}

	if fmt.BitsPerSample != 16 {
		t.Fatalf("Bits per sample is invalid: %d", fmt.BitsPerSample)
	}

	samples, err := reader.ReadSamples(1)
	if len(samples) != 1 {
		t.Fatalf("Length of samples is invalid: %d", len(samples))
	}
	
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}

	if len(bytes) != int(reader.WavData.Size)- (4) {
		t.Fatalf("Data size is invalid: %d", len(bytes))
	}
}