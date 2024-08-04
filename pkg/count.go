package pkg

import (
	"bufio"
	"log/slog"
	"os"
)

func DetermineNumberOfBytes(filePath string) int {
	file, err := os.ReadFile(filePath)
	if err != nil {
		slog.Error("Could not read file", "err", err)
	}

	return len(file)
}

func DetermineNumberOfLines(filePath string) int {
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

	return count
}

func DetermineNumberOfWords(filePath string) int {
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

	return count
}

func DetermineNumberOfCharacters(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		slog.Error("Could not read file", "err", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	count := 0

	for scanner.Scan() {
		count++
	}

	return count
}
