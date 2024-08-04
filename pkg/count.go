package pkg

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
)

func DetermineNumberOfBytes(filePath string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		slog.Error("Could not read file", "err", err)
	}

	fmt.Printf("%d %s\n", len(file), filePath)
}

func DetermineNumberOfLines(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		slog.Error("Could not read file", "err", err)
	}

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		scanner.Text()
		count++
	}

	fmt.Printf("%d %s", count, filePath)
}

func DetermineNumberOfWords(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		slog.Error("Could not read file", "err", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	count := 0

	for scanner.Scan() {
		count++
	}

	fmt.Printf("%d %s", count, filePath)
}
