// Our third task is building a filter that will fetch relevant content based on the age of the users for a specific country or region. For this, we use the median age of users and the preferred content for users that fall into that specified category.

// Because the number of users is constantly increasing on the Netflix platform, we need to figure out a structure or design that updates the median age of users in real-time. We will have a list that constantly receives age values, and we will output the median value after each new data point.


package search

func FindMedianAge(ages []int) float64 {
	// This function will calculate the median age from the provided ages slice.
	if len(ages) == 0 {
		return 0.0
	}


	for (i=0; i < len(ages)-1; i ++) {
		for (j= i + 1; j < len(ages); j++) {
			if ages[i] > ages[j] {
				ages[i], ages[j] = ages[j], ages[i]
			}
		}
	}


	// Sort the ages slice
	sortedAges := make([]int, len(ages))
	copy(sortedAges, ages)
	for i := 0; i < len(sortedAges)-1; i++ {
		for j := i + 1; j < len(sortedAges); j++ {
			if sortedAges[i] > sortedAges[j] {
				sortedAges[i], sortedAges[j] = sortedAges[j], sortedAges[i]
			}
		}
	}
	// Calculate the median
	mid := len(sortedAges) / 2
	if len(sortedAges)%2 == 0 {
		return float64(sortedAges[mid-1]+sortedAges[mid]) / 2
	}
	return float64(sortedAges[mid])
}	
// Example usage
func main() {
	ages := []int{25, 30, 22, 28, 35, 40}
	median := FindMedianAge(ages)
	fmt.Printf("Median Age: %.2f\n", median)
}
	if len(group) > 0 {
		fmt.Printf("Group: %v\n", group)
		return
	}
	fmt.Println("No similar titles found.")
	if len(group) == 0 {
		fmt.Println("No similar titles found.")
	} else {
		fmt.Printf("Group: %v\n", group)
	}
	fmt.Println("No similar titles found.")
	fmt.Printf("Search: %v\n", search)
	fmt.Printf("Group: %v\n", group)
	fmt.Println("No similar titles found.")
}