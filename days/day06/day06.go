package day06

import (
	"fmt"
)

func Part01(in []string) {
	fmt.Println("How many characters need to be processed before the first start-of-packet marker is detected?")
	signal := findSignal(in[0], 4)
	fmt.Println(signal)
}

func Part02(in []string) {
	fmt.Println("How many characters need to be processed before the first start-of-message marker is detected?")
	signal := findSignal(in[0], 14)
	fmt.Println(signal)
}

func findSignal(in string, sliceLen int) int {
	inLen := len(in)
	for i := sliceLen - 1; i < inLen; i++ {
		inMap := make(map[string]int)
		var maxValue int
		for j := i - sliceLen + 1; j <= i; j++ {
			inMap[string(in[j])]++
			if maxValue < inMap[string(in[j])] {
				maxValue = inMap[string(in[j])]
			}
		}

		if maxValue == 1 {
			return i + 1
		}
	}
	return 0
}
