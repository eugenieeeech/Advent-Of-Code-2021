package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Part1()  {
	file, err := os.Open("./input/d2.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	scanner.Split(bufio.ScanWords)

	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Println(input[1])
}
func Task(){
	Part1()
}


