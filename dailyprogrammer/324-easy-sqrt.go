/*
 References:
  https://medium.com/i-math/how-to-find-square-roots-by-hand-f3f7cadf94bb
  https://www.reddit.com/r/dailyprogrammer/comments/6nstip/20170717_challenge_324_easy_manual_square_root/
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func guessSolution(diff func(int) int, min, max int) (int, int, error) {
	if max < min {
		return 0, 0, fmt.Errorf("Max less than min")
	}

	guess := min
	var comp int
	for min <= max {
		mid := (min + max) / 2

		comp = diff(mid)
		if comp < 0 {
			max = mid - 1
		} else if comp > 0 {
			guess = mid
			min = mid + 1
		} else {
			return mid, comp, nil
		}
		fmt.Println("|", comp, mid, guess)
	}

	comp = diff(guess)
	return guess, comp, nil
}

func sqrt(integer, fraction string, precision int) {
	if len(integer)%2 != 0 {
		integer = "0" + integer
	}

	if len(fraction)%2 != 0 {
		fraction = fraction + "0"
	}
	currentPrecision := len(fraction) / 2
	if currentPrecision < precision {
		for ii := 0; ii < precision-currentPrecision; ii++ {
			fraction += "00"
		}
	} else if currentPrecision > precision {
		fraction = fraction[:precision*2]
	}

	integerSteps := len(integer)
	integerIndex := 0
	var A, diff, number int
	var err error
	if integerSteps != 0 {
		number, _ = strconv.Atoi(integer[:2])
		A, diff, err = guessSolution(func(a int) int {
			return number - a*a
		}, 0, 10)
		if err != nil {
			panic(err)
		}
		integerIndex = 2

		var solution int
		for ; integerIndex < len(integer); integerIndex += 2 {
			number, _ = strconv.Atoi(integer[integerIndex : integerIndex+2])
			number = diff*100 + number
			solution, diff, err = guessSolution(func(b int) int {
				return number - b*(b+20*A)
			}, 0, 10)
			if err != nil {
				panic(err)
			}
			A = A*10 + solution
		}
	}

	fractionSteps := len(fraction)
	fractionIndex := 0
	if fractionSteps != 0 {
		if integerIndex == 0 {
			number, _ = strconv.Atoi(fraction[:2])
			A, diff, err = guessSolution(func(a int) int {
				return number - a*a
			}, 0, 10)
			if err != nil {
				panic(err)
			}
			fractionIndex = 2
		}

		var solution int
		for ; fractionIndex < len(fraction); fractionIndex += 2 {
			number, _ = strconv.Atoi(fraction[fractionIndex : fractionIndex+2])
			number = diff*100 + number

			fmt.Println(number, diff, A)
			solution, diff, err = guessSolution(func(b int) int {
				return number - b*(b+20*A)
			}, 0, 10)
			if err != nil {
				panic(err)
			}
			A = A*10 + solution
		}
	}

	fmt.Printf("= %v.%v\n", A, fractionIndex/2)

}

func main() {
	re := regexp.MustCompile(`(\d+) (\d*)\.?(\d*)`)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		if matches == nil || matches[2] == "" && matches[3] == "" {
			fmt.Println("Usage: <precision> <number>")
		} else {
			precision, _ := strconv.Atoi(matches[1])
			sqrt(matches[2], matches[3], precision)
		}
		fmt.Print("> ")
	}
}
