package day07

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const deleteThreshold = 100000
const filesystemSize = 70000000
const updateSize = 30000000

type node struct {
	name      string
	nodeType  string
	size      int
	sonSize   int
	firstNode bool
	sons      []*node
	father    *node
}

func Part01(in []string) {
	fmt.Println("Find all of the directories with a total size of at most 100000. What is the sum of the total sizes of those directories?")

	generateCommandMap(in)
	firstNode := generateFileTree(in)
	var total int
	deepSearchSum(firstNode, &total)
	fmt.Println(total)

}

func Part02(in []string) {
	fmt.Println("Find the smallest directory that, if deleted, would free up enough space on the filesystem to run the update. What is the total size of that directory?")
	generateCommandMap(in)
	firstNode := generateFileTree(in)
	var total int
	deepSearchSum(firstNode, &total)

	freeSpace := filesystemSize - firstNode.size - firstNode.sonSize
	neededSpace := updateSize - freeSpace
	sizeList := make([]int, 0)
	sizeMap := make(map[int]string, len(in))
	sizeList, sizeMap = getFileSizeMap(firstNode, sizeList, sizeMap)
	sort.Ints(sizeList)

	for _, s := range sizeList {
		if s >= neededSpace {
			fmt.Printf("Needed space is %d, so we have to delete dir %s with size %d \n", neededSpace, sizeMap[s], s)
			break
		}
	}
}

func generateCommandMap(in []string) ([]string, map[string][]string) {
	var commandList []string
	var actualCommand string
	commandMap := make(map[string][]string)
	for _, c := range in {
		if string(c[0]) == "$" {
			commandList = append(commandList, c)
			actualCommand = c
		} else {
			commandMap[actualCommand] = append(commandMap[actualCommand], c)
		}
	}

	return commandList, commandMap
}

func generateFileTree(in []string) *node {
	inLen := len(in)
	ft := &node{
		name:      "/",
		nodeType:  "d",
		size:      0,
		sons:      make([]*node, 0),
		father:    nil,
		firstNode: true,
	}

	pointer := ft
	firstNode := ft

	for i, c := range in {
		if strings.Contains(c, "$ cd ") {
			goToFolder := c[5:]
			if !strings.Contains(goToFolder, "/") {
				if strings.Contains(goToFolder, "..") {
					pointer = pointer.father
				} else {
					next, err := findNextFolder(goToFolder, pointer.sons)
					if err == nil {
						pointer = next
					} else {
						fmt.Println("more sad")
					}
				}
			}
		}

		if strings.Contains(c, "$ ls") {
			for j := i + 1; j < inLen && !strings.Contains(in[j], "$"); j++ {
				split := strings.Split(in[j], " ")
				if strings.Contains(split[0], "dir") {
					newNode := &node{
						name:     split[1],
						nodeType: "d",
						father:   pointer,
					}

					pointer.sons = append(pointer.sons, newNode)

				} else {
					split[0] = strings.TrimSpace(split[0])
					size, err := strconv.Atoi(split[0])
					if err == nil {
						pointer.size += size
					}
				}

			}
		}

	}

	return firstNode

}

func deepSearchSum(pointer *node, totalSun *int) (sonSum int) {
	if len(pointer.sons) == 0 {
		if pointer.size <= deleteThreshold {
			*totalSun += pointer.size
		}
		return pointer.size
	}

	for _, p := range pointer.sons {
		sonSize := deepSearchSum(p, totalSun)
		pointer.sonSize += sonSize
	}

	nodeTotal := pointer.sonSize + pointer.size
	if nodeTotal <= deleteThreshold {
		*totalSun += nodeTotal
	}

	return nodeTotal

}

func getFileSizeMap(pointer *node, sizeList []int, sizeMap map[int]string) ([]int, map[int]string) {
	nodeTotal := pointer.size + pointer.sonSize
	sizeList = append(sizeList, nodeTotal)
	sizeMap[nodeTotal] = pointer.name

	if len(pointer.sons) == 0 {
		return sizeList, sizeMap
	}

	for _, p := range pointer.sons {
		sizeList, sizeMap = getFileSizeMap(p, sizeList, sizeMap)
	}

	return sizeList, sizeMap

}

func findNextFolder(name string, nextNodes []*node) (*node, error) {
	for _, f := range nextNodes {
		if f.name == name {
			return f, nil
		}
	}
	return nil, errors.New("sad life")
}
