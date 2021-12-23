package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input/day1.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v\n", err)
	}

	lines := strings.Split(string(b), "\n")
	depth1, err := strconv.Atoi(lines[0])
	if err != nil {
		log.Fatalf("Error parsing line %d: %v\n", 0, err)
	}

	increases := 0
	for i := 1; i < len(lines); i++ {
		depth2, err := strconv.Atoi(lines[i])
		if err != nil {
			log.Fatalf("Error parsing line %d: %v\n", i, err)
		}

		if depth2 > depth1 {
			increases++
		}

		depth1 = depth2
	}

	log.Println("Number of increases:")
	log.Println(increases)
}
