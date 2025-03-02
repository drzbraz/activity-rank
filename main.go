package main

import (
	"container/heap"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Repository represents a repository with its activity score
type Repository struct {
	Name  string
	Score int
}

// MinHeap implements a heap for top repositories
type MinHeap []Repository

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Score < h[j].Score } // Min-heap
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Repository))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// ScoreCalculator calculates repository activity scores
type ScoreCalculator struct{}

func (s ScoreCalculator) Calculate(filesChanged, additions, deletions int) int {
	return (filesChanged * 2) + additions + deletions
}

// CSVLoader loads data from a CSV file
type CSVLoader struct{}

func (c *CSVLoader) LoadCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

// ProcessRepositories calculates scores and keeps the top 10 in a heap
func ProcessRepositories(data [][]string, calculator ScoreCalculator) []Repository {
	repoScores := make(map[string]int)
	for _, row := range data[1:] { // Ignore header
		repo := row[2]
		filesChanged, _ := strconv.Atoi(row[3])
		additions, _ := strconv.Atoi(row[4])
		deletions, _ := strconv.Atoi(row[5])

		repoScores[repo] += calculator.Calculate(filesChanged, additions, deletions)
	}

	// Use a heap for top 10 selection
	h := &MinHeap{}
	heap.Init(h)

	for name, score := range repoScores {
		heap.Push(h, Repository{Name: name, Score: score})
		if h.Len() > 10 {
			heap.Pop(h) // Remove lowest score
		}
	}

	// Convert heap to slice (sorted in descending order)
	topRepos := make([]Repository, h.Len())
	for i := len(topRepos) - 1; i >= 0; i-- {
		topRepos[i] = heap.Pop(h).(Repository)
	}

	return topRepos
}

func main() {
	csvLoader := CSVLoader{}
	data, err := csvLoader.LoadCSV("commits.csv")
	if err != nil {
		fmt.Println("Error loading CSV:", err)
		return
	}

	calculator := ScoreCalculator{}
	topRepos := ProcessRepositories(data, calculator)

	fmt.Println("Top 10 most active repositories:")
	for i, repo := range topRepos {
		fmt.Printf("%d. %s - Score: %d\n", i+1, repo.Name, repo.Score)
	}
}
