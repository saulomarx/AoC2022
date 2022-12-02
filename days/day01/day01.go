package day01

import (
	"fmt"
	"sort"
	"strconv"
)

func Part01(input []string) {
	fmt.Println("Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?")
	var elvesCalories int
	var maxCalories int
	for _, c := range input {
		value, err := strconv.Atoi(c)

		if err == nil {
			elvesCalories += value
		}

		if err != nil {
			if maxCalories < elvesCalories {
				maxCalories = elvesCalories
			}
			elvesCalories = 0
		}

	}

	fmt.Println(maxCalories)

}

func Part02(input []string) {
	fmt.Println("Find the top three Elves carrying the most Calories. How many Calories are those Elves carrying in total?")
	var elvesCalories int
	sumCaloriesByElf := make([]int, 0)
	for _, c := range input {
		value, err := strconv.Atoi(c)

		if err == nil {
			elvesCalories += value
		}

		if err != nil {
			sumCaloriesByElf = append(sumCaloriesByElf, elvesCalories)
			elvesCalories = 0
		}

	}

	sort.Ints(sumCaloriesByElf)
	p := len(sumCaloriesByElf)
	sumTopThree := sumCaloriesByElf[p-1] + sumCaloriesByElf[p-2] + sumCaloriesByElf[p-3]
	fmt.Println(sumTopThree)
}
