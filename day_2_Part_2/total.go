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
	file, err := os.Open("input1.txt")
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return
	}
	
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalFeet := 0

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

		// Calculate volume in cubic feet 
		volume :=  l*w*h
        perim1 := 2*(l+w) 
		perim2 := 2*(w+h)
		perim3 := 2 *(l+h)
		small := perim1
        if perim2 < small {
			small = perim2
		}
		if perim3 < perim2 {
			small = perim3
		}
	

		totalFeet += volume + small
	}

	if err := scanner.Err();
	 err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Total square feet of wrapping paper needed: %d\n", totalFeet)
}

