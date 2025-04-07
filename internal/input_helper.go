package internal

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

func InputHelper(prompt string, shouldHide bool) (string, error) {
	fmt.Print(prompt)

	if shouldHide {
		valueBytes, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return "", fmt.Errorf("InputHelper: ReadPassword: %w", err)
		}
		fmt.Print("\n")
		return string(valueBytes), nil
	}

	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("InputHelper: ReadString: %w", err)
	}
	return value, nil
}
