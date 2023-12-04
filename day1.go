package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	inputFile, err := os.Open("day1data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	// Scanner for the file
	scanner := bufio.NewScanner(inputFile)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		twoDigitNumber, err := extractFirstAndLastDigit(line)
		if err != nil {
			log.Println(err)
			continue
		}
		sum += twoDigitNumber
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

func extractFirstAndLastDigit(s string) (int, error) {
	numbers := regexp.MustCompile(`\d+`).FindAllString(s, -1)
	if len(numbers) == 0 {
		return 0, fmt.Errorf("no numbers found in string")
	}

	firstNumberStr := numbers[0]
	lastNumberStr := numbers[len(numbers)-1]

	firstDigit := firstNumberStr[0]
	lastDigit := lastNumberStr[len(lastNumberStr)-1]

	twoDigitNumberStr := string(firstDigit) + string(lastDigit)
	twoDigitNumber, err := strconv.Atoi(twoDigitNumberStr)
	if err != nil {
		return 0, fmt.Errorf("error converting string to int: %v", err)
	}

	return twoDigitNumber, nil
}
