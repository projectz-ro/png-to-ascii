package utils

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func ClearScreen(except int) {
	if except == 0 {
		fmt.Print("\033[H\033[2J")
	} else {
		fmt.Printf("\033[%d;0H", except+1)
		fmt.Print("\033[J")
	}
}

func PromptWithDefault(scanner *bufio.Scanner, prompt, defaultValue string) string {
	fmt.Printf("%s\n(default: %s): ", prompt, defaultValue)

	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatalf("Error reading input: %v", err)
		}
		return defaultValue
	}

	input := strings.TrimSpace(scanner.Text())
	if input == "" {
		return defaultValue
	}
	return input
}
