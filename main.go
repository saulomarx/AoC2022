package main

import (
	"fmt"

	"github.com/saulomarx/AoC2022/days/day01"
	"github.com/saulomarx/AoC2022/days/day02"
	"github.com/saulomarx/AoC2022/days/day03"
	"github.com/saulomarx/AoC2022/days/day04"
	"github.com/saulomarx/AoC2022/days/day05"
	"github.com/saulomarx/AoC2022/utils"
)

func main() {
	workingDay := 5

	switch workingDay {
	case 1:
		runsDay01()
	case 2:
		runsDay02()
	case 3:
		runsDay03()
	case 4:
		runsDay04()
	case 5:
		runsDay05()
	default:
		fmt.Println("not today")
	}

}

func runsDay01() {
	p := "./days/day01/inputs/in01.txt"
	input := utils.ReadLines(p)
	day01.Part01(input)
	day01.Part02(input)
}

func runsDay02() {
	p := "./days/day02/inputs/in01.txt"
	input := utils.ReadLines(p)
	day02.Part01(input)
	day02.Part02(input)
}

func runsDay03() {
	p := "./days/day03/inputs/in01.txt"
	input := utils.ReadLines(p)
	day03.Part01(input)
	day03.Part02(input)
}

func runsDay04() {
	p := "./days/day04/inputs/in01.txt"
	input := utils.ReadLines(p)
	day04.Part01(input)
	day04.Part02(input)
}

func runsDay05() {
	p := "./days/day05/inputs/in01.txt"
	input := utils.ReadLinesNoTrim(p)
	day05.Part01(input)
	day05.Part02(input)
}
