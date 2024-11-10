package linreg

import (
	"math"

	"01.gritlab.ax/git/mamberla/guess-it-2/ms"
)

// getXValues put the indices of a data set on a dataset of its own
func getXValues(data []float64) (xs []float64) {
	for i := range data {
		xs = append(xs, float64(i))
	}
	return
}

// SlopeAndIntercept returns the angle of a linear regression line  and its position
// on the y axis from a data set
func SlopeAndIntercept(data []float64) (float64, float64) {
	dataXs := getXValues(data)
	meanX := ms.Average(dataXs)
	meanY := ms.Average(data)

	sumDiffsXY := 0.0
	sumDiffsXX := 0.0

	for x, y := range data {
		xFlo := float64(x)
		sumDiffsXY += (xFlo - meanX) * (y - meanY)
		sumDiffsXX += (xFlo - meanX) * (xFlo - meanX)
	}

	slope := sumDiffsXY / sumDiffsXX
	interc := intercept(meanX, meanY, slope)

	return slope, interc
}

// SlopeAndIntercept2 returns the angle of a linear regression line  and its position
// on the y axis from a data set
func SlopeAndIntercept2(data []float64) (float64, float64) {
	slope := Pearson(data) * ms.StandardDeviation(data) / ms.StandardDeviation(getXValues(data))
	interc := intercept(ms.Average(getXValues(data)), ms.Average(data), slope)
	return slope, interc
}

// Intercept returns the location of a linear regression line on the y-axis, based on the means
// of the x and y values (mx, my) and the line's slope (slope)
func intercept(mx, my, slope float64) (itc float64) {
	return my - slope*mx
}

// Pearson returns the Pearson correlation coefficient of a dataset, where
// x value is the index of a y value
func Pearson(d []float64) float64 {
	sumX := 0.0  // sum of x values
	sumY := 0.0  // sum of y values
	sumXX := 0.0 // sum of squared x values
	sumYY := 0.0 // sum of squared y values
	sumXY := 0.0 // sum of cross products

	for x, y := range d {
		xFlo := float64(x)

		sumX += xFlo
		sumY += y

		sumXX += xFlo * xFlo
		sumYY += y * y

		sumXY += xFlo * y
	}

	dividend := float64(len(d))*sumXY - sumX*sumY
	divisor := math.Sqrt((float64(len(d))*sumXX - sumX*sumX) * (float64(len(d))*sumYY - sumY*sumY))

	return dividend / divisor
}
