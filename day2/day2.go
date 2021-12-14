/*
--- Day 2: Dive! ---
*/
package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Part1()  {
	var temp int
	horizontal := 0
	depth := 0
	file, err := os.Open("./input/d2.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	scanner.Split(bufio.ScanWords)

	//var input []string

	for scanner.Scan() {
		switch scanner.Text(){
			case "forward":
				scanner.Scan()
				temp,err = strconv.Atoi(scanner.Text())
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				horizontal+= temp
			case "up":
				scanner.Scan()
				temp,err = strconv.Atoi(scanner.Text())
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				depth -= temp
			case "down":
				scanner.Scan()
				temp,err = strconv.Atoi(scanner.Text())
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				depth += temp
		}
	}

	fmt.Println(horizontal * depth)
}

func Part2()  {
	var aim int
	var temp int
	horizontal := 0
	depth := 0
	file, err := os.Open("./input/d2.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	scanner.Split(bufio.ScanWords)

	//var input []string

	for scanner.Scan() {
		switch scanner.Text(){
		case "forward":
			scanner.Scan()
			temp,err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			horizontal+= temp
			depth += aim*temp
		case "up":
			scanner.Scan()
			temp,err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			aim -= temp
		case "down":
			scanner.Scan()
			temp,err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			aim += temp
		}
	}

	fmt.Println(horizontal * depth)
}

func Task(){
	Part1()
	Part2()
}


