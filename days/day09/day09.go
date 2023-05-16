package day09

import (
	"fmt"
	"strconv"
	"strings"
)

type tuple struct {
	x int
	y int
}

func Part01(in []string) {
	fmt.Println("Simulate your complete hypothetical series of motions. How many positions does the tail of the rope visit at least once?")
	talesMap := make(map[tuple]bool)
	actualPlace := tuple{}
	hPlace := tuple{}
	tPlace := tuple{}

	talesMap[actualPlace] = true

	for _, l := range in {
		d, qtd := getDirection(l)
		fmt.Println(d, qtd)
		switch d {
		case "R":
			hPlace, tPlace = walkR(qtd, hPlace, tPlace, talesMap)
		case "L":
			hPlace, tPlace = walkL(qtd, hPlace, tPlace, talesMap)
		case "U":
			hPlace, tPlace = walkU(qtd, hPlace, tPlace, talesMap)
		case "D":
			hPlace, tPlace = walkD(qtd, hPlace, tPlace, talesMap)
		}
	}

	// var count int
	// for k, v := range talesMap {
	// 	if v {
	// 		count++
	// 	}
	// 	fmt.Println(k)

	// }
	// fmt.Println(count, len(talesMap))
}

func Part02(in []string) {
	fmt.Println("Consider each tree on your map. What is the highest scenic score possible for any tree?")
}

func walkR(qtd int, hPlace, tPlace tuple, talesMap map[tuple]bool) (tuple, tuple) {
	for i := 0; i < qtd; i++ {
		if i == 0 {
			hPlace.x++
			talesMap[tPlace] = true
		} else {
			if tPlace.y != hPlace.y {
				tPlace.y = hPlace.y
			}
			tPlace.x++
			hPlace.x++
			talesMap[tPlace] = true
		}
		fmt.Printf("\nhPlace %d,%d === tPlace %d,%d\n", hPlace.x, hPlace.y, tPlace.x, tPlace.y)

	}
	return hPlace, tPlace
}

func walkL(qtd int, hPlace, tPlace tuple, talesMap map[tuple]bool) (tuple, tuple) {
	for i := 0; i < qtd; i++ {
		if i == 0 {
			hPlace.x--
			talesMap[tPlace] = true
		} else {
			if tPlace.y != hPlace.y {
				tPlace.y = hPlace.y
			}
			tPlace.x--
			hPlace.x--
			talesMap[tPlace] = true
		}
		fmt.Printf("\nhPlace %d,%d === tPlace %d,%d\n", hPlace.x, hPlace.y, tPlace.x, tPlace.y)
	}
	return hPlace, tPlace
}

func walkU(qtd int, hPlace, tPlace tuple, talesMap map[tuple]bool) (tuple, tuple) {
	for i := 0; i < qtd; i++ {
		if i == 0 {
			hPlace.y++
			talesMap[tPlace] = true
		} else {
			if tPlace.x != hPlace.x {
				tPlace.x = hPlace.x
			}
			tPlace.y++
			hPlace.y++
			talesMap[tPlace] = true
		}
		fmt.Printf("\nhPlace %d,%d === tPlace %d,%d\n", hPlace.x, hPlace.y, tPlace.x, tPlace.y)

	}
	return hPlace, tPlace
}

func walkD(qtd int, hPlace, tPlace tuple, talesMap map[tuple]bool) (tuple, tuple) {
	for i := 0; i < qtd; i++ {
		if i == 0 {
			hPlace.y--
			talesMap[tPlace] = true
		} else {
			if tPlace.x != hPlace.x {
				tPlace.x = hPlace.x
			}
			tPlace.y--
			hPlace.y--
			talesMap[tPlace] = true
		}
		fmt.Printf("\nhPlace %d,%d === tPlace %d,%d\n", hPlace.x, hPlace.y, tPlace.x, tPlace.y)

	}
	return hPlace, tPlace
}

func getDirection(l string) (d string, qtd int) {
	split := strings.Split(l, " ")
	value, _ := strconv.Atoi(split[1])

	return split[0], value

}
