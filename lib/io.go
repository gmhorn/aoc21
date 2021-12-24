package lib

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadLinesToInts(path string) ([]int, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return nil, err
	}

	data := make([]int, 0)
	for i, line := range lines {
		val, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("could not parse line %d in %s: %v\n", i, path, err)
		}
		data = append(data, val)
	}

	return data, nil
}

func ReadLines(path string) ([]string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(b), "\n"), nil
}
