package day02

import "fmt"

const r = "ROCK"
const p = "PAPER"
const s = "SCISSOR"

const w = "WIN"
const l = "LOSE"
const d = "DRAW"

var winPoints = map[string]int{r: 1, p: 2, s: 3}
var strategy = map[string]string{"X": l, "Y": d, "Z": w}
var winStrategy = map[string]string{p: s, s: r, r: p}
var loseStrategy = map[string]string{s: p, r: s, p: r}
var playMap = map[string]string{"A": r, "X": r, "B": p, "Y": p, "C": s, "Z": s}

func Part01(in []string) {
	fmt.Println("What would your total score be if everything goes exactly according to your strategy guide?")
	var roundPoints int
	for _, game := range in {
		o := playMap[string(game[0])]
		m := playMap[string(game[2])]
		roundPoints = roundPoints + resolveGame(o, m)
	}
	fmt.Println(roundPoints)

}

func Part02(in []string) {
	fmt.Println("Following the Elf's instructions for the second column, what would your total score be if everything goes exactly according to your strategy guide?")
	var roundPoints int
	for _, game := range in {
		o := string(game[0])
		m := string(game[2])

		o, m = changePlay(o, m)

		roundPoints = roundPoints + resolveGame(o, m)
	}
	fmt.Println(roundPoints)

}

func resolveGame(o, m string) int {
	roundPoints := winPoints[m]

	if o != m {
		// I play rock
		if m == r {
			if o == p {
				return roundPoints
			}
			return roundPoints + 6
		}

		// I play Scissor
		if m == s {
			if o == r {
				return roundPoints
			}
			return roundPoints + 6
		}

		// I play paper
		if m == p {
			if o == s {
				return roundPoints
			}
			return roundPoints + 6
		}
	}

	// Same play
	return roundPoints + 3
}

func changePlay(o, m string) (string, string) {
	result := strategy[m]

	o = playMap[o]

	if result == w {
		m = winStrategy[o]
	}

	if result == l {
		m = loseStrategy[o]
	}

	if result == d {
		m = o
	}

	return o, m

}
