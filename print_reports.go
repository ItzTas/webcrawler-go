package main

import (
	"fmt"
	"sort"
	"strconv"
)

func printReports(pages map[string]int, baseURL string) {
	fmt.Printf(`
=============================
  REPORT for %s
=============================
`, baseURL)

	fmt.Println("")
	for _, sorted := range sortReports(pages) {
		fmt.Printf("Found %s internal links to %s\n", sorted[1], sorted[0])
	}
}

func sortReports(pages map[string]int) [][]string {
	values := make([][]string, 0, len(pages))

	for k, v := range pages {
		values = append(values, []string{k, strconv.Itoa(v)})
	}

	sort.Slice(values, func(i, j int) bool {
		val1, _ := strconv.Atoi(values[i][1])
		val2, _ := strconv.Atoi(values[j][1])
		return val1 > val2
	})

	return values
}
