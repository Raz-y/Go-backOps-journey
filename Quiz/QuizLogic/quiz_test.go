package quiz

import (
	"testing"
)

func TestParseLines(t *testing.T) {
	lines := [][]string{
		{"question1", " answer1 "},
		{"question2 ", "answer2"},
	}
	expected := []problem{
		{q: "question1", a: "answer1"},
		{q: "question2", a: "answer2"},
	}

	problems := parseLines(lines)
	if len(problems) != len(expected) {
		t.Fatalf("Expected %d problems, but got %d", len(expected), len(problems))
	}

	for i, problem := range problems {
		if problem.q != expected[i].q || problem.a != expected[i].a {
			t.Errorf("Expected problem %v, but got %v", expected[i], problem)
		}
	}
}

func TestCheckAnswer(t *testing.T) {
	tests := []struct {
		answer   string
		expected string
		want     bool
	}{
		{" answer1 ", "answer1", true},
		{" answer2 ", "answer1", false},
		{" Answer ", "answer1", false},
		{" answer1 ", " answer1", true},
	}

	for _, tt := range tests {
		if got := checkAnswer(tt.answer, tt.expected); got != tt.want {
			t.Errorf("checkAnswer(%q, %q) = %v; want %v", tt.answer, tt.expected, got, tt.want)
		}
	}
}
