package response

// BodyWordsegPost represents the request body for the /wordseg endpoint
type BodyWordsegPost struct {
	Text string `json:"text" validate:"required"`
}

// GuessResult represents the structure of a guess result
type GuessResult struct {
	Slot   int    `json:"slot"`
	Guess  string `json:"guess"`
	Result string `json:"result"` // Values: "absent", "present", "correct"
}

// ValidationError represents a single validation error
type ValidationError struct {
	Loc  []string `json:"loc"`
	Msg  string   `json:"msg"`
	Type string   `json:"type"`
}

// HTTPValidationError represents the structure of a validation error response
type HTTPValidationError struct {
	Detail []ValidationError `json:"detail"`
}

type LetterFeedback struct {
	Letter string `json:"letter"`
	Status string `json:"status"` // "correct", "present", "absent"
}

type GuessResponse struct {
	Feedback []LetterFeedback `json:"feedback"`
	Message  string           `json:"message"`
}
