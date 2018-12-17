package main

import (
	"fmt"
	"os"
)

const (
	deptIncrement = 4
)

func main() {
	paths := os.Args[1:]
	if len(paths) == 0 {
		paths = []string{"."}

	}

	var dirTotal, fileTotal uint
	depth := 0
	for _, path := range paths {
		// fmt.Printf("%*v\n", depth, path)
		dirCount, fileCount := processDir(path, depth)

		dirTotal += dirCount
		fileTotal += fileCount
	}
	fmt.Printf("%v directories, %v files\n", dirTotal, fileTotal)
}

func processDir(path string, depth int) (dirCount uint, fileCount uint) {
	dir, err := os.Open(path)
	if err != nil {
		fmt.Printf("%v [Error opening dir]\n", path)
		return
	}
	defer dir.Close()

	entries, err := dir.Readdir(0)
	if err != nil {
		return
	}

	// proces each entry in this dir
	lastIndex := len(entries) - 1
	// s := "├── "
	for i, fi := range entries {
		if i == lastIndex {
			// s = "└── "
		}
		// fmt.Printf("%*v%v\n", depth, s, fi.Name())
		if !fi.IsDir() {
			fileCount++
			continue
		}

		dirCount++
		newPath := fmt.Sprintf("%v%c%v", path, os.PathSeparator, fi.Name())
		dCount, fCount := processDir(newPath, depth+deptIncrement)
		fileCount += fCount
		dirCount += dCount
	}
	return
}
