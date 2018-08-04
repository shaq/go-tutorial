package main

import "fmt"

func main() {
	grades := make(map[string]float32)

	// Assign k-v pairs in grades map.
	grades["Shaquille"] = 73
	grades["Jo"] = 90
	fmt.Println(grades)

	// Get from value from map using given key.
	shaqsGrade := grades["Shaquille"]
	fmt.Println("Shaquille's grade:", shaqsGrade)

	// Delete given key from map.
	delete(grades, "Shaquille")
	fmt.Println(grades)

	// Iterate through map.
	for k, v := range grades {
		fmt.Println(k, ":", v)
	}
}
