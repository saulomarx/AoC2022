package day05

import (
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	qtd  int
	from int
	to   int
}

func Part01(in []string) {
	fmt.Println("After the rearrangement procedure completes, what crate ends up on top of each stack?")
	piles, instructions := getPilesAndInstructions(in)

	mappedInstruction := mountInstructions(instructions)

	n := getNumberPiles(piles[len(piles)-1])
	pileMap := mountPiles(piles)

	println("First Pile")
	printPiles(pileMap, n)
	for _, i := range mappedInstruction {
		movePiles(pileMap, i.qtd, i.from, i.to)

	}

	println("========")
	println("Moved Pile")
	printPiles(pileMap, n)

	println("========")
	println("Last crates")
	printTopCrates(pileMap, n)
}

func Part02(in []string) {
	fmt.Println("After the rearrangement procedure completes, what crate ends up on top of each stack?")
	piles, instructions := getPilesAndInstructions(in)

	mappedInstruction := mountInstructions(instructions)

	n := getNumberPiles(piles[len(piles)-1])
	pileMap := mountPiles(piles)

	println("First Pile")
	printPiles(pileMap, n)
	for _, i := range mappedInstruction {
		movePiles9001(pileMap, i.qtd, i.from, i.to)

	}

	println("========")
	println("Moved Pile")
	printPiles(pileMap, n)

	println("========")
	println("Last crates")
	printTopCrates(pileMap, n)
}

func mountInstructions(in []string) []instruction {
	instructions := make([]instruction, len(in))
	for idx, stringInstruction := range in {
		var newInstruction instruction
		stringInstruction = stringInstruction[4:]

		splitFrom := strings.Split(stringInstruction, "from")

		move := strings.TrimSpace(splitFrom[0])
		newInstruction.qtd, _ = strconv.Atoi(move)

		splitTo := strings.Split(splitFrom[1], "to")
		splitTo[0] = strings.TrimSpace(splitTo[0])
		splitTo[1] = strings.TrimSpace(splitTo[1])

		newInstruction.from, _ = strconv.Atoi(splitTo[0])
		newInstruction.to, _ = strconv.Atoi(splitTo[1])

		instructions[idx] = newInstruction
	}

	return instructions
}

func getPilesAndInstructions(in []string) (piles, instructions []string) {
	for i, line := range in {
		if len(line) >= 4 && line[:4] == "move" {
			return in[:i-1], in[i:]
		}
	}

	return
}

func getNumberPiles(s string) int {
	s = strings.TrimSpace(s)
	lastPile := string(s[len(s)-1])

	iP, _ := strconv.Atoi(lastPile)

	return iP
}

func mountPiles(piles []string) map[int][]string {
	pileMap := make(map[int][]string)
	for i := len(piles) - 2; i >= 0; i-- {
		lineLen := len(piles[i])
		for j, pileNumber := 1, 1; j <= lineLen; j, pileNumber = j+4, pileNumber+1 {

			box := string(piles[i][j])
			if strings.TrimSpace(box) != "" {
				pileMap[pileNumber] = append(pileMap[pileNumber], box)
			}
		}
	}
	return pileMap
}

func movePiles(piles map[int][]string, qtd, from, to int) {
	auxMove := len(piles[from]) - qtd
	aux := piles[from][auxMove:]
	aux = reverseStringSlice(aux)

	piles[from] = piles[from][:auxMove]
	piles[to] = append(piles[to], aux...)
}

func movePiles9001(piles map[int][]string, qtd, from, to int) {
	auxMove := len(piles[from]) - qtd
	aux := piles[from][auxMove:]

	piles[from] = piles[from][:auxMove]
	piles[to] = append(piles[to], aux...)
}

func printPiles(piles map[int][]string, numberOfPiles int) {

	for i := 1; i <= numberOfPiles; i++ {
		fmt.Println(i, piles[i])
	}
}

func printTopCrates(piles map[int][]string, numberOfPiles int) {

	for i := 1; i <= numberOfPiles; i++ {
		pileSize := len(piles[i])
		fmt.Printf(piles[i][pileSize-1])
	}
	fmt.Println()
}

func reverseStringSlice(o []string) []string {
	oLen := len(o)

	if oLen <= 1 {
		return o
	}

	n := make([]string, oLen)
	copy(n, o)

	for i := oLen/2 - 1; i >= 0; i-- {
		opp := oLen - i - 1
		n[i], n[opp] = n[opp], n[i]
	}

	return n
}
