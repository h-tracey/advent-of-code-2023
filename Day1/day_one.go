package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func getNums(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return 10
}

func ConvertNamedNums(str string) string {
	numNames := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for k, v := range numNames {
		str = replaceAllNames(str, k, fmt.Sprint(v))
	}
	return str
}

func replaceAllNames(s string, substr string, replace string) string {
	if strings.Contains(s, substr) {
		idx := strings.Index(s, substr)
		replaced := s[:idx+1] + replace + s[idx+1:]
		return replaceAllNames(replaced, substr, replace)

	} else {
		return s
	}
}
func getFirst(s string) (string, error) {
	for _, letter := range s {
		first := getNums(string(letter))
		if first < 10 {
			return string(letter), nil
		}
	}
	return "", errors.New("no num")
}

func firstLastInts(s string) int {
	converted := ConvertNamedNums(s)
	firstNum, _ := getFirst(converted)
	lastNum, _ := getFirst(reverse(converted))
	combined := firstNum + lastNum
	return getNums(combined)

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	calibration := 0
	for scanner.Scan() {
		calibration += firstLastInts(scanner.Text())
	}
	fmt.Println(calibration)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
