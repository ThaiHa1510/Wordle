package utils

import "Wordle/internal/response"

// CompareWords compares the guess to the target word and returns feedback for each letter
func CompareWords(guess, target string) []response.LetterFeedback {
	feedback := make([]response.LetterFeedback, len(guess))
	targetRunes := []rune(target)
	guessRunes := []rune(guess)

	matched := make([]bool, len(targetRunes))

	for i := 0; i < len(guessRunes); i++ {
		if guessRunes[i] == targetRunes[i] {
			feedback[i] = response.LetterFeedback{
				Letter: string(guessRunes[i]),
				Status: "correct",
			}
			matched[i] = true
		}
	}

	for i := 0; i < len(guessRunes); i++ {
		if feedback[i].Status == "correct" {
			continue
		}
		found := false
		for j := 0; j < len(targetRunes); j++ {
			if !matched[j] && guessRunes[i] == targetRunes[j] {
				found = true
				matched[j] = true
				break
			}
		}
		if found {
			feedback[i] = response.LetterFeedback{
				Letter: string(guessRunes[i]),
				Status: "present",
			}
		} else {
			feedback[i] = response.LetterFeedback{
				Letter: string(guessRunes[i]),
				Status: "absent",
			}
		}
	}

	return feedback
}
