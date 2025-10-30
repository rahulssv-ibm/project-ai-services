package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readUserInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return response, nil
}

func ConfirmAction() (bool, error) {
	response, err := readUserInput()
	if err != nil {
		return false, fmt.Errorf("failed to take user input: %w", err)
	}

	response = strings.TrimSpace(strings.ToLower(response))

	// If response is 'n' or 'no' -> do not proceed with deletion
	if response == "n" || response == "no" {
		return false, nil
	}

	// if response is neither 'y' and 'yes' -> then its an invalid input
	if response != "y" && response != "yes" {
		return false, fmt.Errorf("received invalid input: %s. Please respond with 'y' or 'n'", response)
	}
	return true, nil
}
