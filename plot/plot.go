package plot

import (
	"os"
	"os/exec"
	"log"
)

//Spectrogram runs python code to generate spectrogram image of a track
func Spectrogram(filename string, id string) {
	//And now for the tricky bit...
	cmd := exec.Command("python3", "../plot/plot.py", filename, id)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr	

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
