package day08

import (
	"fmt"

	"github.com/saulomarx/AoC2022/utils"
)

func Part01(in []string) {
	fmt.Println("Consider your map; how many trees are visible from outside the grid?")

	mtx := utils.ReadIntMatrix(in)
	checkMtx := createCheckMtx(len(mtx), len(mtx[0]))

	lineSize := len(mtx)
	columnSize := len(mtx[0])

	var sum int
	for i := 0; i < lineSize; i++ {
		sum += leftView(i, mtx, checkMtx)
		sum += rightView(i, mtx, checkMtx)
	}

	for j := 0; j < columnSize; j++ {
		sum += topView(j, mtx, checkMtx)
		sum += bottomView(j, mtx, checkMtx)
	}

	fmt.Println(sum)

}

func Part02(in []string) {
	fmt.Println("Consider each tree on your map. What is the highest scenic score possible for any tree?")
	mtx := utils.ReadIntMatrix(in)

	lineSize := len(mtx)
	columnSize := len(mtx[0])

	var max int
	for i := 0; i < lineSize; i++ {
		for j := 0; j < columnSize; j++ {
			u := lookUp(i, j, mtx)
			d := lookDown(i, j, mtx)
			l := lookLeft(i, j, mtx)
			r := lookRight(i, j, mtx)

			mult := u * d * l * r

			if mult > max {
				max = mult
			}
		}
	}

	fmt.Println(max)

}

func leftView(i int, mtx [][]int, checkMtx [][]bool) int {
	maxSize := mtx[i][0]
	var count int

	for j := range mtx[i] {
		if j == 0 && !checkMtx[i][j] {
			checkMtx[i][j] = true
			count++
		}
		if mtx[i][j] > maxSize {
			if !checkMtx[i][j] {
				checkMtx[i][j] = true
				count++
			}
			maxSize = mtx[i][j]
		}
	}
	return count
}

func rightView(i int, mtx [][]int, checkMtx [][]bool) int {
	maxSize := mtx[i][len(mtx[i])-1]
	var count int

	for j := len(mtx[i]) - 1; j >= 0; j-- {
		if j == len(mtx[i])-1 && !checkMtx[i][j] {
			checkMtx[i][j] = true
			count++
		}
		if mtx[i][j] > maxSize {
			if !checkMtx[i][j] {
				checkMtx[i][j] = true
				count++
			}
			maxSize = mtx[i][j]
		}
	}
	return count
}

func topView(j int, mtx [][]int, checkMtx [][]bool) int {
	maxSize := mtx[0][j]
	var count int

	for i := range mtx {
		if i == 0 && !checkMtx[i][j] {
			count++
			checkMtx[i][j] = true
		}
		if mtx[i][j] > maxSize {
			if !checkMtx[i][j] {
				checkMtx[i][j] = true
				count++
			}
			maxSize = mtx[i][j]
		}
	}
	return count
}

func bottomView(j int, mtx [][]int, checkMtx [][]bool) int {
	maxSize := mtx[len(mtx)-1][j]
	var count int

	for i := len(mtx) - 1; i >= 0; i-- {
		if i == len(mtx)-1 && !checkMtx[i][j] {
			checkMtx[i][j] = true
			count++
		}
		if mtx[i][j] > maxSize {
			if !checkMtx[i][j] {
				checkMtx[i][j] = true
				count++
			}
			maxSize = mtx[i][j]
		}
	}

	return count
}

func lookUp(i, j int, mtx [][]int) int {
	var tree int
	for k := i - 1; k >= 0; k-- {
		tree++
		if mtx[k][j] >= mtx[i][j] {
			return tree
		}
	}
	return tree
}

func lookDown(i, j int, mtx [][]int) int {
	var tree int
	lineLen := len(mtx)
	for k := i + 1; k < lineLen; k++ {
		tree++
		if mtx[k][j] >= mtx[i][j] {
			return tree
		}
	}
	return tree
}

func lookRight(i, j int, mtx [][]int) int {
	var tree int
	for k := j - 1; k >= 0; k-- {
		tree++
		if mtx[i][k] >= mtx[i][j] {
			return tree
		}
	}
	return tree
}

func lookLeft(i, j int, mtx [][]int) int {
	var tree int
	columnSize := len(mtx[0])
	for k := j + 1; k < columnSize; k++ {
		tree++

		if mtx[i][k] >= mtx[i][j] {
			return tree
		}
	}
	return tree
}

func createCheckMtx(i, j int) [][]bool {
	mtx := make([][]bool, i)
	for k := range mtx {
		mtx[k] = make([]bool, j)
	}
	return mtx
}
