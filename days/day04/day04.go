package day04

import (
	"fmt"
	"strconv"
	"strings"
)

type cleanerRange struct {
	begin int
	end   int
}

func Part01(in []string) {
	fmt.Println("In how many assignment pairs does one range fully contain the other?")

	var sumOverlaps int
	for _, cr := range in {
		elf01, elf02 := createsCleanerRange(cr)
		contains := checkFullOverlap(elf01, elf02)
		if contains {
			sumOverlaps++
		}
	}
	fmt.Println(sumOverlaps)
}

func Part02(in []string) {
	fmt.Println("In how many assignment pairs do the ranges overlap?")
	var sumOverlaps int
	for _, cr := range in {
		elf01, elf02 := createsCleanerRange(cr)
		contains := checkOverlap(elf01, elf02)
		if contains {
			sumOverlaps++
		}
	}
	fmt.Println(sumOverlaps)
}

func createsCleanerRange(in string) (elf01, elf02 cleanerRange) {
	ranges := strings.Split(in, ",")

	first := strings.Split(ranges[0], "-")
	second := strings.Split(ranges[1], "-")

	elf01.begin, _ = strconv.Atoi(first[0])
	elf01.end, _ = strconv.Atoi(first[1])

	elf02.begin, _ = strconv.Atoi(second[0])
	elf02.end, _ = strconv.Atoi(second[1])

	return
}

func checkFullOverlap(elf01, elf02 cleanerRange) bool {
	if elf01.begin <= elf02.begin && elf01.end >= elf02.end {
		return true
	}

	if elf02.begin <= elf01.begin && elf02.end >= elf01.end {
		return true
	}

	return false
}

func checkOverlap(elf01, elf02 cleanerRange) bool {
	if elf01.begin <= elf02.begin && elf01.end >= elf02.begin {
		return true
	}

	if elf01.begin <= elf02.end && elf01.end >= elf02.end {
		return true
	}

	if elf02.begin <= elf01.begin && elf02.end >= elf01.begin {
		return true
	}

	if elf02.begin <= elf01.end && elf02.end >= elf01.end {
		return true
	}

	return false
}
