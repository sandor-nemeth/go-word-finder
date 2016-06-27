package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/sandor-nemeth/levenshtein"
)

func main() {
	f := "/usr/share/dict/words"
	lines := make([]string, 235886, 235886)

	readDictionary(f, lines)

	start := time.Now()

	dist, _ := strconv.Atoi(os.Args[2])

	simple(os.Args[1], int32(dist), lines)

	elapsed := time.Since(start)
	log.Printf("Search took %s", elapsed)
}

func simple(in string, d int32, lines []string) {
	for _, w := range lines {
		cost := slev.SLev(in, w)

		if cost <= d {
			fmt.Printf("%s %d\n", w, cost)
		}
	}
}

func readDictionary(path string, lines []string) {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	var i int = 0
	for scanner.Scan() {
		lines[i] = scanner.Text()
		i++
	}
}
