package main

import (
	"fmt"
	"testing"
)

// write test cases for the function
func TestGroupTitles(t *testing.T) {
	titles := []string{"pushpa", "puspha", "adipurush", "Adipurush", "apuhsp", "kubera"}
	search := "puspha"

	groups := groupTitles(titles)
	searchVector := vectorizeTitle(search)

	group := groups[fmt.Sprintf("%v", searchVector)]
	fmt.Printf("Group: %v", group)

	if len(group) != 3 {
		t.Errorf("Expected 2 titles in group, got %d", len(group))
	}

	if group[0] != "pushpa" || group[1] != "puspha" || group[2] != "apuhsp" {
		t.Errorf("Expected pushpa and puspha in group, got %v", group)
	}
}