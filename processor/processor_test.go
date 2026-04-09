package processor

import (
	"testing"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "hex conversion",
			input:    "1E (hex) files were added",
			expected: "30 files were added",
		},
		{
			name:     "binary conversion",
			input:    "It has been 10 (bin) years",
			expected: "It has been 2 years",
		},
		{
			name:     "uppercase",
			input:    "Ready, set, go (up) !",
			expected: "Ready, set, GO!",
		},
		{
			name:     "lowercase",
			input:    "I should stop SHOUTING (low)",
			expected: "I should stop shouting",
		},
		{
			name:     "capitalize",
			input:    "welcome to the island (cap)",
			expected: "welcome to the Island",
		},
		{
			name:     "multiple uppercase",
			input:    "this is so exciting (up, 2)",
			expected: "this is SO EXCITING",
		},
		{
			name:     "punctuation spacing",
			input:    "Hello , there ! How are you ? I am fine .",
			expected: "Hello, there! How are you? I am fine.",
		},
		{
			name:     "punctuation exceptions",
			input:    "Wait... what !?",
			expected: "Wait... what!?",
		},
		{
			name:     "quotes",
			input:    "As turning said , ' I think therefore I am '",
			expected: "As turning said, 'I think therefore I am'",
		},
		{
			name:     "articles",
			input:    "There is it a hour a egg an dog",
			expected: "There is it an hour an egg a dog",
		},
		{
			name:     "complex test from prompt",
			input:    "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			expected: "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Process(tt.input)
			if result != tt.expected {
				t.Errorf("Process() = %q, expected %q", result, tt.expected)
			}
		})
	}
}
