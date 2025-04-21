package main

import (
	"fmt"
	"math"
	"math/rand"
)

// ========================================================Question 1====================================================
// Provided random function: returns 0 or 1 with 50% probability
func random() int {
	return rand.Intn(2)
}

// Custom function: returns 0 with 75% and 1 with 25% probability
func customRandom() int {
	for {
		// Generate two random bits
		bit1 := random()
		bit2 := random()

		// Create a number in the range 0–3
		num := bit1*2 + bit2

		// Use 0, 1, 2 for output 0 (75%), and 3 for output 1 (25%)
		if num == 3 {
			return 1
		} else {
			return 0
		}
	}
}

// ===================================================End of Logic Question 1================================================

// ========================================================Question 2====================================================
func secondHighestPaidSalary(salaries []int) int {
	if len(salaries) < 2 {
		panic("Need at least two salaries")
	}

	max1 := math.MinInt64
	max2 := math.MinInt64

	for _, salary := range salaries {
		if salary > max1 {
			max2 = max1
			max1 = salary
		} else if salary > max2 && salary != max1 {
			max2 = salary
		}
	}

	if max2 == math.MinInt64 {
		panic("No second highest salary found (all values may be the same)")
	}

	return max2
}

// ===================================================End of Logic Question 2================================================

// ========================================================Question 3====================================================
func mostFrequentDigit(str string) int {
	freq := make([]int, 10) // For digits 0–9

	// Count each digit's frequency
	for _, ch := range str {
		freq[ch-'0']++
	}

	maxFreq := -1
	result := 0
	for digit, count := range freq {
		if count > maxFreq {
			maxFreq = count
			result = digit
		}
	}

	return result
}

// ===================================================End of Logic Question 3================================================
// ========================================================Question 4====================================================
func elementExists(element string, array []string) bool {
	left, right := 0, len(array)-1

	for left <= right {
		mid := left + (right-left)/2
		if array[mid] == element {
			return true
		} else if array[mid] < element {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// ===================================================End of Logic Question 4================================================
// ========================================================Question 5====================================================
func findHighestAltitude(gain []int) int {
	altitude := 0
	maxAltitude := 0

	for _, g := range gain {
		altitude += g
		if altitude > maxAltitude {
			maxAltitude = altitude
		}
	}

	return maxAltitude
}

// ===================================================End of Logic Question 5================================================

func main() {
	// ========================================================Question 1====================================================
	count0, count1 := 0, 0
	for i := 0; i < 100000; i++ {
		if customRandom() == 0 {
			count0++
		} else {
			count1++
		}
	}
	fmt.Printf("0: %d, 1: %d\n", count0, count1)
	// =================================================End of Question 1====================================================

	// ========================================================Question 2====================================================
	salaries := []int{5000, 3000, 6000, 1000}
	fmt.Println("Second highest salary is:", secondHighestPaidSalary(salaries))
	// =================================================End of Question 2====================================================

	// ========================================================Question 3====================================================
	input := "19827439111498912111"
	fmt.Println("Most frequent digit is:", mostFrequentDigit(input))
	// =================================================End of Question 3====================================================
	// ========================================================Question 4====================================================
	array := []string{"apple", "boy", "cat", "dog", "elephant"}
	element := "car"
	fmt.Println("Is Element Exists?", elementExists(element, array))
	// =================================================End of Question 4====================================================
	// ========================================================Question 5====================================================
	gain := []int{-5, 1, 5, 0, -7}
	fmt.Println("Highest altitude:", findHighestAltitude(gain))
	// =================================================End of Question 5====================================================
}
