package plot

import (
	"fmt"
	"math"
	//"os"
	"log"

	"gonum.org/v1/gonum/mat"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette"
	"gonum.org/v1/plot/plotter"
)

type Grid struct {
	c, r int

	Data mat.Matrix
}

func (g Grid) Dims() (c, r int) {
	r, c = g.Data.Dims() 
	
	return c, r
}

func (g Grid) X(c int) float64 {
	_, n := g.Data.Dims()
	if c < 0 || c >= n {
		panic("index out of range")
	}

	return float64(c)
}

func (g Grid) Y(r int) float64 {
	m, _ := g.Data.Dims()
	if r < 0 || r >= m {
		panic("index out of range")
	}

	return float64(r)
}

func (g Grid) Z(c, r int) float64 {
	return g.Data.At(r, c)
}

type stringTicks struct {
	Labels []string
}

func (s stringTicks) Ticks(min, max float64) []plot.Tick {
	var t []plot.Tick
	
	for i := math.Trunc(min); i <= max; i++ {
		t = append(t, plot.Tick{Value: i, Label: s.Labels[int(i)]})
	}
	
	return t
}

// PlotSpectrogram - Plots spectrogram for a given file and saves it as an image
func PlotSpectrogram(values [][]float64) {
	dat := make([]float64, 0)

	for row, rows := range values {
		for col, _ := range rows {
			dat = append(dat, values[row][col])
		}
	}

	m := Grid{
	    Data: mat.NewDense(1, 2048, dat),
	}

	pal := palette.Heat(100, 1)
	h := plotter.NewHeatMap(m, pal)
	
	p, err := plot.New()
	if err != nil {
	    log.Panic(err)
	}

	//Rasterize "pixels" to get an image
	//h.Rasterized = true

	p.Add(h)

	err = p.Save(750, 500, "heatmap.png")
	if err != nil {
		log.Panic(err)
	}
}
