package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"01.gritlab.ax/git/mamberla/guess-it-2/guessing"
	"01.gritlab.ax/git/mamberla/guess-it-2/ms"
)

// isOutlier tells if a data point is not like the others
func isOutlier(n float64, nums []float64) bool {
	return ms.Abs(n-ms.Average(nums)) > ms.Variance(nums)*0.17
}

// readInputAndGuess expects numerical values from the standard input and guesses
// a range for the next number
func readInputAndGuess() {
	scanner := bufio.NewScanner(os.Stdin)
	numbers := []float64{}
	added := 0
	all := 0
	r1, r2 := 0, 0

	for scanner.Scan() {
		txt := scanner.Text()
		num, err := strconv.ParseFloat(txt, 64)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		all++

		// If outlier: use old guess and ignore the input
		if added > 6 && isOutlier(num, numbers) {
			fmt.Printf("%v %v\n", r1, r2)
			continue
		}

		// Data set max 61 long
		if added < 61 {
			added++
			numbers = append(numbers, num)
		} else {
			numbers = append(numbers[1:], num)
		}

		r1, r2 = guessing.Linear(numbers)
		//r1, r2 = all+104, all+196 // Wins datasets 4 and 5 every time

		fmt.Printf("%v %v\n", r1, r2)
	}
}

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: ./guess-it-2")
		os.Exit(1)
	}
	readInputAndGuess()
}
