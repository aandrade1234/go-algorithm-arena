package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(alienOrder([]string{"wrt", "wrf", "er", "ett", "rftt"}))
	fmt.Println(alienOrder([]string{"z", "x"}))
	fmt.Println(alienOrder([]string{"z", "x", "z"}))
}

func alienOrder(words []string) string {
	graph := map[byte][]byte{}
	indegree := map[byte]int{}

	for _, word := range words {
		for i := range word {
			indegree[word[i]] = 0
		}
	}

	// Create the graph
	for i := range len(words) - 1 {
		word1 := words[i]
		word2 := words[i+1]
		minLength := min(len(word1), len(word2))

		if len(word1) > len(word2) && strings.HasPrefix(word1, word2) {
			return ""
		}

		for j := range minLength {
			if word1[j] != word2[j] {
				graph[word1[j]] = append(graph[word1[j]], word2[j])
				indegree[word2[j]]++

				break
			}
		}
	}

	// Execute the BFS
	var result strings.Builder

	queue := []byte{}

	// Prepare the queue
	for char, count := range indegree {
		if count == 0 {
			queue = append(queue, char)
		}
	}

	// process the queue
	for len(queue) > 0 {
		char := queue[0]
		queue = queue[1:]

		result.WriteByte(char)

		for _, neighbor := range graph[char] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if result.Len() != len(indegree) {
		return ""
	}

	return result.String()
}
