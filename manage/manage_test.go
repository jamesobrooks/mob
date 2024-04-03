package manage

import (
	"errors"
	"testing"
)

func TestGetNextTypist(t *testing.T) {
	tests := []struct {
		name          string
		currentTypist string
		participants  []string
		expected      string
		expectedError error
	}{
		{
			name:          "Next participant exists",
			currentTypist: "Joe",
			participants:  []string{"Joe", "Jane", "Jack"},
			expected:      "Jane",
			expectedError: nil,
		},
		{
			name:          "Current participant is last, loop back",
			currentTypist: "Jack",
			participants:  []string{"Joe", "Jane", "Jack"},
			expected:      "Joe",
			expectedError: nil,
		},
		{
			name:          "Current participant not found",
			currentTypist: "Jill",
			participants:  []string{"Joe", "Jane", "Jack"},
			expected:      "",
			expectedError: errors.New("current timer user not found"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := getNextTypist(tc.currentTypist, tc.participants)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
			if (err != nil && tc.expectedError == nil) || (err == nil && tc.expectedError != nil) {
				t.Errorf("Expected error %v, got %v", tc.expectedError, err)
			}
			if err != nil && tc.expectedError != nil && err.Error() != tc.expectedError.Error() {
				t.Errorf("Expected error message '%v', got '%v'", tc.expectedError.Error(), err.Error())
			}
		})
	}
}
