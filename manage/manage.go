package manage

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
)

const fileName = "participants.txt"

func AddParticipants(participants []string) {
	existing := readParticipants()

	existingMap := make(map[string]bool)
	for _, participant := range existing {
		existingMap[participant] = true
	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	for _, newParticipant := range participants {
		if _, exists := existingMap[newParticipant]; !exists {
			if _, err := file.WriteString(newParticipant + "\n"); err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}
	}
}

func RemoveParticipants(participants []string) {
	contents, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(contents), "\n")
	var newLines []string

	for _, line := range lines {
		if line != "" && !slices.Contains(participants, line) {
			newLines = append(newLines, line)
		}
	}

	newContent := strings.Join(newLines, "\n")
	newContent = strings.TrimSpace(newContent) + "\n"

	err = os.WriteFile(fileName, []byte(newContent), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}

func DisplayNextTypist(currentTypist string) {
	participants := readParticipants()

	if len(participants) == 0 {
		fmt.Println("No participants found.")
		return
	}

	nextTypist, err := getNextTypist(currentTypist, participants)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Next typist:", nextTypist)
}

func readParticipants() []string {
	file, err := os.Open(fileName)
	if err != nil {
		return []string{}
	}
	defer file.Close()

	var participants []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		participants = append(participants, scanner.Text())
	}
	//// Closing explicitly because tests seem to get old contents
	//file.Close()

	return participants
}

func getNextTypist(currentTypist string, participants []string) (string, error) {
	currentPosition := slices.Index(participants, currentTypist)
	if currentPosition == -1 {
		return "", errors.New("current timer user not found")
	}
	if currentPosition < len(participants)-1 {
		return participants[currentPosition+1], nil
	} else {
		return participants[0], nil
	}
}
