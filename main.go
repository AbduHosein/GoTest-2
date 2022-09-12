package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func benchmarkParallelsum(nums []int, ch chan int) {
	// Initialize variables to store random numbers and sum
	sum := 0

	for _, v := range nums {
		sum += v
	}

	ch <- sum
}

func benchmarkSlicestable(x int) {
	var s []int
	now := time.Now()

	for i := 0; i < x; i++ {
		randomNum := rand.Intn(100)
		s = append(s, randomNum)

	}

	// Sorting slice using anonymous function -- stack overflow
	start := now
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	elapsed := time.Since(start)
	fmt.Println(elapsed)

	// Sorting slice using StableSLice
	start = now
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	elapsed = time.Since(start)
	fmt.Println(elapsed)

}

func main() {
	// Initialize variable to be used in functions
	var x int
	var nums []int
	now := time.Now()
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	// Prompt user for integer value, handle nil values --stack overflow
	fmt.Println("Enter an integer value:")
	_, err := fmt.Scanf("%d", &x)
	if err != nil {
		fmt.Println(err)

	}

	// Generate x random numbers and store in nums array
	for i := 0; i < x; i++ {
		randomNum := rand.Intn(100)
		nums = append(nums, randomNum)

	}

	// Benchmarking sum calculation without go routines
	fmt.Println("Performance of parallel summation:")
	start := now
	go benchmarkParallelsum(nums, c1)
	elapsed := time.Since(start)
	fmt.Println(<-c1)
	fmt.Println(elapsed)

	// Benchmarking sum with routines -- unsure why I am getting a routine asleep error but in theory this should work :/
	start = now
	go benchmarkParallelsum(nums[:len(nums)/2], c2)
	go benchmarkParallelsum(nums[len(nums)/2:], c3)
	sum1, sum2 := <-c2, <-c3
	elapsed = time.Since(start)
	fmt.Println(sum1 + sum2)
	fmt.Println(elapsed)

	// Benchmarking sorting slices with stable sort -- trying for extra credit
	fmt.Println("Performance of sorting slices:")
	benchmarkSlicestable(x)
}
