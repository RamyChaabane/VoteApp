package vote_test

import (
	"testing"

	"github.com/RamyChaabane/VoteApp/backend/internal/domain/vote"
	"github.com/stretchr/testify/assert"
)

func TestIsValidOption(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{name: "valid cats", input: "Cats", expected: true},
		{name: "valid dogs", input: "Dogs", expected: true},
		{name: "invalid lowercase", input: "cats", expected: false},
		{name: "invalid other", input: "Birds", expected: false},
		{name: "empty string", input: "", expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, vote.IsValidOption(tt.input))
		})
	}
}
