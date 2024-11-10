package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"01.gritlab.ax/git/mamberla/guess-it-2/guessing"
	"01.gritlab.ax/git/mamberla/guess-it-2/ms"
)

// guessNextRange makes a guess about in which range the next number will be
func guessNextRange(nums []int) (int, int) {

	if len(nums) == 0 {
		return 0, 0
	}
	if len(nums) == 1 {
		return int(nums[0] - 10), int(nums[0] + 10)
	}

	numsF := toFloats(nums)
	rng := guessing.Linear(numsF)

	return rng[0], rng[1]
}

// isOutlier tells if a data point is not like the others
func isOutlier(n int, nums []int) bool {
	floats := toFloats(nums)
	return ms.Abs(float64(n)-ms.Average(floats)) > ms.Variance(floats)*0.17
}

// toFloats converts a slice of ints to float64s
func toFloats(nums []int) []float64 {
	floats := make([]float64, len(nums))
	for i, n := range nums {
		floats[i] = float64(n)
	}
	return floats
}

func main() {
	fmt.Println(os.Args)

	if len(os.Args) != 1 {
		fmt.Println("Usage: ./guess-it-2")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	numbers := []int{}
	seen := 0
	r1, r2 := 0, 0

	for scanner.Scan() {
		txt := scanner.Text()
		num, err := strconv.Atoi(txt)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		// Use old guess and ignore the input if it's an outlier
		if seen > 6 && isOutlier(num, numbers) {
			fmt.Printf("%v %v\n", r1+1, r2+1)
			continue
		}

		// Use a data set max 61 long
		if seen < 61 {
			seen++
			numbers = append(numbers, num)
		} else {
			numbers = append(numbers[1:], num)
		}

		r1, r2 = guessNextRange(numbers)
		fmt.Printf("%v %v\n", r1, r2)
	}
}
