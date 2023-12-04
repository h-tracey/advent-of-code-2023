package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getNums(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return -1
}

func ArrayContains(array []string) bool {
	specialChars := []string{"#", "*", "$", "-", "%", "+", "@", "&", "=", "/"}
	for _, item := range array {
		if slices.Contains(specialChars, item) {
			return true
		}
	}
	return false
}

func subArrayGearFinder(array []string) int {
	for i, item := range array {
		if item == "*" {
			return i
		}
	}
	return -1
}

func treatAsArrays(input string) []string {
	var out []string
	splitter := func(r rune) bool {
		return strings.ContainsRune(".#*$-%+@&=/", r)
	}
	charBuffer := ""
	lastN := len(input) - 1
	for i, char := range input {
		if splitter(char) && charBuffer == "" {
			out = append(out, string(char))
		} else if splitter(char) {

			out = append(out, charBuffer)
			out = append(out, string(char))
			charBuffer = ""
		} else {
			charBuffer = charBuffer + string(char)
			if i == lastN {
				out = append(out, charBuffer)
				charBuffer = ""
			}
		}
	}
	return out
}

func schematicNumber(row []string, high []string, low []string) bool {
	if ArrayContains(row) {
		return false
	}
	if ArrayContains(high) {
		return false
	}
	if ArrayContains(low) {
		return false
	}
	return true
}

func getPaddedSearchSpace(index int, span int, cellRow []string, higher []string, lower []string) ([]string, []string, []string, int, int) {
	checkEnd := index + span + 1
	checkStart := index - 1

	if checkEnd > len(cellRow) {
		checkEnd--
	}

	if checkStart == -1 {
		checkStart++
	}
	if len(lower) != 0 {
		lower = lower[checkStart:checkEnd]
	}
	if len(higher) != 0 {
		higher = higher[checkStart:checkEnd]
	}
	cellRow = cellRow[checkStart:checkEnd]

	return cellRow, higher, lower, checkStart, checkEnd
}
func checkSafety(index int, span int, cellRow []string, higher []string, lower []string) bool {
	searchSpaceC, searchSpaceH, searchSpaceL, _, _ := getPaddedSearchSpace(index, span, cellRow, higher, lower)
	return schematicNumber(searchSpaceC, searchSpaceH, searchSpaceL)

}

func getGearIndex(searchSpace []string, index int) int {
	gearCell := subArrayGearFinder(searchSpace)
	if gearCell >= 0 {
		return gearCell + index + 1
	}
	return -1

}

func checkGears(row, index int, span int, cellRow []string, higher []string, lower []string) (int, int) {
	searchSpaceC, searchSpaceH, searchSpaceL, iStart, _ := getPaddedSearchSpace(index, span, cellRow, higher, lower)
	c := getGearIndex(searchSpaceC, iStart)
	if c >= 0 {
		return c, row
	}
	h := getGearIndex(searchSpaceH, iStart)
	if h >= 0 {
		return h, row - 1
	}
	l := getGearIndex(searchSpaceL, iStart)
	if l >= 0 {
		return l, row + 1
	}
	return -1, -1

}

func sum(nums []int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

type GearPos struct {
	row, column int
}
type Gears struct {
	number int
	gear   GearPos
}

func getGearRatio(gears []Gears) {
	var ratios [][]int
	var gearLoc []GearPos
	for _, gear := range gears {
		if len(gearLoc) == 0 {
			gearLoc = append(gearLoc, gear.gear)
			ratio := []int{gear.number}
			ratios = append(ratios, ratio)
		} else {
			appended := false
			for i, loc := range gearLoc {
				if gear.gear.column == loc.column && gear.gear.row == loc.row {
					ratios[i] = append(ratios[i], gear.number)
					appended = true
				}
			}
			if !appended {
				gearLoc = append(gearLoc, gear.gear)
				ratio := []int{gear.number}
				ratios = append(ratios, ratio)

			}
		}
	}
	var ratioGears []int
	for _, ratio := range ratios {
		if len(ratio) > 1 {
			ratioGears = append(ratioGears, ratio[0]*ratio[1])
		}
	}
	powers := 0
	for _, gearPower := range ratioGears {
		powers += gearPower
	}
	fmt.Println(powers)
}

func checkAllArrays(numericArrays [][]string, allArrays [][]string) int {
	lastRow := len(allArrays) - 1
	var validNums []int
	var validGears []Gears
	for i, row := range numericArrays {
		maskedIndex := 0
		for _, cell := range row {
			numValue := getNums(cell)
			if numValue > 0 {
				var higher []string
				var lower []string
				if i != 0 {
					higher = allArrays[i-1]
				}
				if i != lastRow {
					lower = allArrays[i+1]
				}
				result := checkSafety(maskedIndex, len(cell), allArrays[i], higher, lower)
				if !result {
					validNums = append(validNums, numValue)
				}
				gear, gearRow := checkGears(i, maskedIndex, len(cell), allArrays[i], higher, lower)
				if gear >= 0 {
					validGears = append(validGears, Gears{number: numValue, gear: GearPos{row: gearRow, column: gear}})

				}
				maskedIndex += len(cell)

			} else {
				maskedIndex++
			}

		}
	}
	getGearRatio(validGears)
	return sum(validNums)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var schematicArrays []string
	for scanner.Scan() {
		schematicArrays = append(schematicArrays, scanner.Text())
	}
	var treatedArrays [][]string
	var arrayMask [][]string
	for _, schamaticLine := range schematicArrays {
		treatedArrays = append(treatedArrays, treatAsArrays(schamaticLine))
		arrayMask = append(arrayMask, strings.Split(schamaticLine, ""))

	}
	fmt.Println(checkAllArrays(treatedArrays, arrayMask))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

//p.1 = 539713
