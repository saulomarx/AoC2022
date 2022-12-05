package day03

import "fmt"

func Part01(in []string) {
	fmt.Println("Find the item type that appears in both compartments of each rucksack. What is the sum of the priorities of those item types?")
	var prioritySum int

	for _, l := range in {
		prioritySum += findRepeatedPriority(l)
	}

	fmt.Println(prioritySum)
}

func findRepeatedPriority(s string) int {
	size := len(s)
	half := size / 2
	firstCompartment := make(map[rune]int, size)
	for i := range s[:half] {
		l := rune(s[i])
		firstCompartment[l] = letterToPriority(l)
	}

	for _, c := range s[half:] {
		if p, ok := firstCompartment[c]; ok {
			return p
		}

	}
	return 0
}

func Part02(in []string) {
	fmt.Println("Find the item type that corresponds to the badges of each three-Elf group. What is the sum of the priorities of those item types?")
	var prioritySum int
	inLen := len(in)
	for i := 0; i < inLen; i += 3 {
		c := findBadges(in[i], in[i+1], in[i+2])
		p := letterToPriority(c)
		prioritySum += p
	}

	fmt.Println(prioritySum)
}

func findBadges(f, s, t string) rune {
	firstBag := make(map[rune]bool, len(f))
	secondBag := make(map[rune]bool, len(s))

	for _, c := range f {
		firstBag[c] = true
	}

	for _, c := range s {
		secondBag[c] = true
	}

	for _, c := range t {
		_, okFirst := firstBag[c]
		_, okSecond := secondBag[c]

		if okFirst == okSecond && okFirst {
			return c
		}
	}

	return rune(0)

}

func letterToPriority(l rune) int {
	if l == 0 {
		return 0
	}

	if l >= 97 {
		return int(l) - 96
	}

	return int(l) - 38
}
