package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ReadLines(path string) []string {
	var input []string
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Deu ruim")
	}
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)

		input = append(input, line)
	}

	return input
}

func ReadLinesNoTrim(path string) []string {
	var input []string
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Deu ruim")
	}
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = line[:len(line)-1]
		input = append(input, line)
	}

	return input
}

func ReadIntMatrix(in []string) [][]int {
	mtx := make([][]int, len(in))
	for i, l := range in {
		newLine := make([]int, len(l))
		for j, c := range l {
			value, _ := strconv.Atoi(string(c))
			newLine[j] = value
		}
		mtx[i] = newLine
	}
	return mtx
}

func PrintIntMtx(in [][]int) {
	for i := range in {
		for j := range in[i] {
			fmt.Printf("%d", in[i][j])
		}
		fmt.Printf("\n")
	}
}

func PrintBollMtx(in [][]bool) {
	for i := range in {
		for j := range in[i] {
			var v int
			if in[i][j] {
				v = 1
			}
			fmt.Printf("%d", v)
		}
		fmt.Printf("\n")
	}
}
