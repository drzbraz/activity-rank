package main

import (
	"testing"
)

// TestScoreCalculator verifies that the score calculation works correctly.
func TestScoreCalculator(t *testing.T) {
	calculator := ScoreCalculator{}

	tests := []struct {
		filesChanged int
		additions    int
		deletions    int
		expected     int
	}{
		{1, 10, 5, 17},
		{2, 20, 10, 34},
		{3, 0, 0, 6},
		{0, 5, 5, 10},
	}

	for _, tt := range tests {
		result := calculator.Calculate(tt.filesChanged, tt.additions, tt.deletions)
		if result != tt.expected {
			t.Errorf("Expected %d, got %d", tt.expected, result)
		}
	}
}

// TestProcessRepositories ensures the top repositories are correctly selected.
func TestProcessRepositories(t *testing.T) {
	data := [][]string{
		{"Commit", "Author", "Repository", "FilesChanged", "Additions", "Deletions"}, // Header
		{"1", "Alice", "repo1", "3", "10", "5"},
		{"2", "Bob", "repo2", "2", "20", "10"},
		{"3", "Charlie", "repo3", "5", "10", "3"},
		{"4", "Alice", "repo1", "1", "5", "5"},
		{"5", "Bob", "repo2", "3", "15", "5"},
		{"6", "Charlie", "repo3", "2", "5", "5"},
	}

	calculator := ScoreCalculator{}
	topRepos := ProcessRepositories(data, calculator)

	// Check that we get the correct number of top repositories
	expectedCount := 3
	if len(topRepos) != expectedCount {
		t.Errorf("Expected %d repositories, got %d", expectedCount, len(topRepos))
	}

	// Verify that repo2 is ranked highest based on the given data
	if topRepos[0].Name != "repo2" {
		t.Errorf("Expected repo2 at the top, got %s", topRepos[0].Name)
	}
}
