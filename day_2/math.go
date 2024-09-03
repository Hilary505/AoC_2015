package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalPaper := 0

	// Read each line from the file
	for scanner.Scan() {
		line := scanner.Text()
		dimensions := strings.Split(line, "x")

		if len(dimensions) != 3 {
			fmt.Println("Invalid dimensions format:", line)
			//continue
		}

		// Parse dimensions
		l, err := strconv.Atoi(dimensions[0])
		if err != nil {
			fmt.Println("Error parsing length:", err)
			//continue
		}
		w, err := strconv.Atoi(dimensions[1])
		if err != nil {
			fmt.Println("Error parsing width:", err)
			//continue
		}
		h, err := strconv.Atoi(dimensions[2])
		if err != nil {
			fmt.Println("Error parsing height:", err)
			//continue
		}

		// Calculate surface area and slack
		surfaceArea := 2*(l*w) + 2*(w*h) + 2*(h*l)

		s1 := l * w
		s2 := w * h
		s3 := l * h
		small := s1

		switch {
		case s2 < small:
			small = s2

		case s3 < small:
			small = s3
		default:

		}

		totalPaper += surfaceArea + small

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Printf("Total square feet of wrapping paper needed: %d\n", totalPaper)
}
