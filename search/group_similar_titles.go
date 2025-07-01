package main

import (
	"fmt"
	"strings"
)

// Create a vector representation of the title
func vectorizeTitle(title string) [26]int {
	v := [26]int{}
	for _, char := range strings.ToLower(title) {
		if char >= 'a' && char <= 'z' {
			v[char-'a']++
		}
	}	
	return v
}

// groupTitles groups titles that are similar to each other
func groupTitles(titles []string) map[string][]string {
	groups := make(map[string][]string)
	
	for _, title := range titles {
		v := vectorizeTitle(title)
		// Convert vector to string key for grouping
		key := fmt.Sprintf("%v", v)
		groups[key] = append(groups[key], title)
	}
	
	return groups
}

//./group_similar_titles.go:44:2: fmt.Println call has possible formatting directive %v
func main() {

	titles := []string{"pushpa", "puspha", "adipurush", "Adipurush", "apuhsp", "kubera"}
	search := "puspha"
	fmt.Printf("Search: %v", search)

	groups := groupTitles(titles)
	
	searchVector := vectorizeTitle(search)

	group := groups[fmt.Sprintf("%v", searchVector)]
	fmt.Printf("Group: %v", group)
}
