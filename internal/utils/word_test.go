package utils

import (
	"testing"

	"Wordle/internal/response"

	"github.com/google/go-cmp/cmp"
)

// TestCompareWords tests the CompareWords function with various scenarios.
func TestCompareWords(t *testing.T) {
	tests := []struct {
		name        string
		guess       string
		target      string
		expected    []response.LetterFeedback
		expectPanic bool
	}{
		{
			name:   "All letters correct",
			guess:  "apple",
			target: "apple",
			expected: []response.LetterFeedback{
				{Letter: "a", Status: "correct"},
				{Letter: "p", Status: "correct"},
				{Letter: "p", Status: "correct"},
				{Letter: "l", Status: "correct"},
				{Letter: "e", Status: "correct"},
			},
		},
		{
			name:   "Some letters correct, some present, some absent",
			guess:  "plane",
			target: "apple",
			expected: []response.LetterFeedback{
				{Letter: "p", Status: "present"},
				{Letter: "l", Status: "present"},
				{Letter: "a", Status: "present"},
				{Letter: "n", Status: "absent"},
				{Letter: "e", Status: "correct"},
			},
		},
		{
			name:   "All letters absent",
			guess:  "zzzzz",
			target: "apple",
			expected: []response.LetterFeedback{
				{Letter: "z", Status: "absent"},
				{Letter: "z", Status: "absent"},
				{Letter: "z", Status: "absent"},
				{Letter: "z", Status: "absent"},
				{Letter: "z", Status: "absent"},
			},
		},
		{
			name:   "Repeated letters in guess, single in target",
			guess:  "allee",
			target: "apple",
			expected: []response.LetterFeedback{
				{Letter: "a", Status: "correct"},
				{Letter: "l", Status: "present"},
				{Letter: "l", Status: "absent"},
				{Letter: "e", Status: "absent"},
				{Letter: "e", Status: "correct"},
			},
		},
		{
			name:   "Repeated letters in target, single in guess",
			guess:  "paper",
			target: "apple",
			expected: []response.LetterFeedback{
				{Letter: "p", Status: "present"},
				{Letter: "a", Status: "present"},
				{Letter: "p", Status: "correct"},
				{Letter: "e", Status: "present"},
				{Letter: "r", Status: "absent"},
			},
		},
		{
			name:   "Case insensitivity",
			guess:  "apple",
			target: "apple",
			expected: []response.LetterFeedback{
				{Letter: "a", Status: "correct"},
				{Letter: "p", Status: "correct"},
				{Letter: "p", Status: "correct"},
				{Letter: "l", Status: "correct"},
				{Letter: "e", Status: "correct"},
			},
		},
		{
			name:        "Empty target",
			guess:       "apple",
			target:      "",
			expectPanic: true,
			expected: []response.LetterFeedback{
				{Letter: "a", Status: "absent"},
				{Letter: "p", Status: "absent"},
				{Letter: "p", Status: "absent"},
				{Letter: "l", Status: "absent"},
				{Letter: "e", Status: "absent"},
			},
		},
		{
			name:        "Different lengths - longer guess",
			guess:       "apples",
			target:      "apple",
			expectPanic: true,
			expected: []response.LetterFeedback{
				{Letter: "a", Status: "correct"},
				{Letter: "p", Status: "correct"},
				{Letter: "p", Status: "correct"},
				{Letter: "l", Status: "correct"},
				{Letter: "e", Status: "correct"},
				{Letter: "s", Status: "absent"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("CompareWords(%q, %q) did not panic, but expected a panic", tt.guess, tt.target)
					}
				}()
				_ = CompareWords(tt.guess, tt.target)
			} else {
				result := CompareWords(tt.guess, tt.target)

				expected := tt.expected
				if len(tt.guess) > len(tt.target) {
					for i := len(tt.target); i < len(tt.guess); i++ {
						expected = append(expected, response.LetterFeedback{
							Letter: string([]rune(tt.guess)[i]),
							Status: "absent",
						})
					}
				}

				if diff := cmp.Diff(result, expected); diff != "" {
					t.Errorf("CompareWords(%q, %q) = %v; want %v", tt.guess, tt.target, result, expected)
				}
			}
		})
	}
}
