package main

import (
	"bufio"
	"math/rand"
	"os"
)

type RainbowTable []string

func loadRainbowTableFromFile(filename string) (table RainbowTable, err error) {
	f, err := os.Open(filename)

	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		table = append(table, scanner.Text())
	}

	err = scanner.Err()

	return
}

func (t RainbowTable) Len() int {
	return len(t)
}

func (t RainbowTable) GetRandomElement() string {
	return t[rand.Intn(len(t))]
}
