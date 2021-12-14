package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Part1()  {
	var count = 0
	file, err := os.Open("./input/d1.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	scanner.Split(bufio.ScanLines)
	var input []int

	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		input = append(input, x)
		if len(input) > 1 {
			if input[len(input)-1] > input[len(input)-2] {
				count++
			}
		}

	}

	fmt.Println(count)
}
func Part2(){
	var count = 0
	file, err := os.Open("input/d1.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	scanner.Split(bufio.ScanLines)
	var input []int

	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		input = append(input, x)
	}
	for i := 0; i < len(input)-3 ; i++ {
		if (input[i+1]+input[i+2]+input[i+3]) > (input[i]+input[i+1]+input[i+2]) {
			count++
		}
	}

	fmt.Println(count)
}

func Task() {
	Part1()
	Part2()
}
