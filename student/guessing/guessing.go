package guessing

import (
	"01.gritlab.ax/git/mamberla/guess-it-2/linreg"
	"01.gritlab.ax/git/mamberla/guess-it-2/ms"
)

func Linear(data []float64) (int, int) {
	slope, intercept := linreg.SlopeAndIntercept(data)

	x := float64(len(data) + 1)
	value := x*slope + intercept
	deviate := 37.5

	return ms.RoundToInt(value - deviate), ms.RoundToInt(value + deviate)
}
