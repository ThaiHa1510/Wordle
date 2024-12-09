package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGetRandomWord(t *testing.T) {
	// Setup: Initialize the wordList with sample data
	wordList = []string{"apple", "banana", "grape", "orange", "berry", "melon"}

	tests := []struct {
		name        string
		size        int
		seed        int64
		expected    string
		expectError bool
	}{
		{
			name:        "Valid size with exact seed",
			size:        6,
			seed:        0,
			expected:    "banana",
			expectError: false,
		},
		{
			name:        "Valid size without seed (random)",
			size:        5,
			seed:        0,  // Indicating random seed
			expected:    "", // Cannot predict; check length
			expectError: false,
		},
		{
			name:        "No words of specified size",
			size:        10,
			seed:        0,
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		tt := tt // Capture range variable
		t.Run(tt.name, func(t *testing.T) {
			word, err := GetRandomWord(tt.size, tt.seed)
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error for size %d, but got none", tt.size)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for size %d: %v", tt.size, err)
				}
				if tt.expected != "" {
					if word != tt.expected {
						t.Errorf("Expected word %q, got %q", tt.expected, word)
					}
				} else {
					if len(word) != tt.size {
						t.Errorf("Expected word length %d, got %d", tt.size, len(word))
					}
				}
			}
		})
	}
}

func TestIsValidWord(t *testing.T) {
	// Setup: Initialize the wordList with sample data
	wordList = []string{"apple", "banana", "grape", "orange", "berry", "melon"}

	tests := []struct {
		name     string
		word     string
		expected bool
	}{
		{
			name:     "Word exists (lowercase)",
			word:     "apple",
			expected: true,
		},
		{
			name:     "Word exists (uppercase)",
			word:     "BANANA",
			expected: true,
		},
		{
			name:     "Word does not exist",
			word:     "kiwi",
			expected: false,
		},
		{
			name:     "Empty string",
			word:     "",
			expected: false,
		},
	}

	for _, tt := range tests {
		tt := tt // Capture range variable
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidWord(tt.word)
			if result != tt.expected {
				t.Errorf("IsValidWord(%q) = %v; want %v", tt.word, result, tt.expected)
			}
		})
	}
}

func TestGetDailyWord(t *testing.T) {
	// Setup: Initialize the dailyList with sample data
	dailyList = []string{"sunny", "cloudy", "rainy", "stormy", "windy"}

	tests := []struct {
		name        string
		size        int
		seed        int64
		expected    string
		expectError bool
	}{
		{
			name:        "Valid size with exact seed",
			size:        5,
			seed:        2,
			expected:    "rainy",
			expectError: false,
		},
		{
			name:        "Valid size without seed (random)",
			size:        9,
			seed:        1,
			expected:    "",
			expectError: true,
		},
		{
			name:        "No daily words of specified size",
			size:        8,
			seed:        1,
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		tt := tt // Capture range variable
		t.Run(tt.name, func(t *testing.T) {
			word, err := GetDailyWord(tt.size, tt.seed)
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error for daily size %d, but got none", tt.size)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for daily size %d: %v", tt.size, err)
				}
				if tt.expected != "" {
					if word != tt.expected {
						t.Errorf("Expected daily word %q, got %q", tt.expected, word)
					}
				} else {
					if len(word) != tt.size {
						t.Errorf("Expected daily word length %d, got %d", tt.size, len(word))
					}
				}
			}
		})
	}
}

func TestAddNewWord(t *testing.T) {
	// Setup: Initialize the wordList with sample data
	wordList = []string{"apple", "banana", "grape"}

	// Create a temporary directory to simulate the file system
	tempDir := t.TempDir()
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}

	// Change the working directory to the temporary directory
	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatalf("Failed to change working directory: %v", err)
	}
	defer func() {
		// Restore the original working directory after the test
		os.Chdir(originalWd)
	}()

	// Create a mock words.txt file in the expected path
	// Considering the AddNewWord constructs the path as "../../internal/utils/words.txt"
	mockPath := filepath.Join(tempDir, "../../internal/utils")
	err = os.MkdirAll(mockPath, os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create mock directory: %v", err)
	}

	wordsFilePath := filepath.Join(mockPath, "words.txt")
	err = ioutil.WriteFile(wordsFilePath, []byte("apple\nbanana\ngrape\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to create mock words.txt: %v", err)
	}

	tests := []struct {
		name        string
		newWord     string
		expected    []string
		expectError bool
	}{
		{
			name:        "Add valid new word",
			newWord:     "kiwi",
			expected:    []string{"apple", "banana", "grape", "kiwi"},
			expectError: false,
		},
		{
			name:        "Add existing word",
			newWord:     "apple",
			expected:    []string{"apple", "banana", "grape"},
			expectError: true,
		},
		{
			name:        "Add empty string",
			newWord:     "",
			expected:    []string{"apple", "banana", "grape"},
			expectError: true,
		},
		{
			name:        "Add word with non-alphabetic characters",
			newWord:     "kiwi123",
			expected:    []string{"apple", "banana", "grape"},
			expectError: true,
		},
	}

	for _, tt := range tests {
		tt := tt // Capture range variable
		t.Run(tt.name, func(t *testing.T) {
			err := AddNewWord(tt.newWord)
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error when adding word %q, but got none", tt.newWord)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error when adding word %q: %v", tt.newWord, err)
				} else {
					// Verify that the word was added to wordList
					if len(wordList) != len(tt.expected) {
						t.Errorf("wordList length = %d; want %d", len(wordList), len(tt.expected))
					}
					for i, word := range tt.expected {
						if wordList[i] != word {
							t.Errorf("wordList[%d] = %q; want %q", i, wordList[i], word)
						}
					}

					// Verify that the word was added to words.txt
					content, err := ioutil.ReadFile(wordsFilePath)
					if err != nil {
						t.Errorf("Failed to read words.txt: %v", err)
					}
					lines := strings.Split(strings.TrimSpace(string(content)), "\n")
					if len(lines) != len(tt.expected) {
						t.Errorf("words.txt line count = %d; want %d", len(lines), len(tt.expected))
					}
					for i, word := range tt.expected {
						if lines[i] != word {
							t.Errorf("words.txt line %d = %q; want %q", i, lines[i], word)
						}
					}
				}
			}
		})
	}
}
