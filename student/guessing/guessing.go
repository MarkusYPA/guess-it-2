package guessing

import (
	"01.gritlab.ax/git/mamberla/guess-it-2/linreg"
	"01.gritlab.ax/git/mamberla/guess-it-2/ms"
)

func Linear(data []float64) [2]int {
	rng := [2]int{}

	slope, intercept := linreg.SlopeAndIntercept(data)

	x := float64(len(data) + 1)
	value := x*slope + intercept
	deviate := 37.5

	rng[0], rng[1] = ms.RoundToInt(value-deviate), ms.RoundToInt(value+deviate)

	return rng
}
