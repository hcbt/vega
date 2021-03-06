package plot

import (
	"os"
	"os/exec"
	"path/filepath"
	"log"
)

//Spectrogram runs python code to generate spectrogram image of a track
func Spectrogram(cwd string, id string) {
	filename := id + ".wav"

	//And now for the tricky bit...
	cmd := exec.Command("python3", filepath.Join(cwd, "../plot/plot.py"), filename, id)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr	

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
