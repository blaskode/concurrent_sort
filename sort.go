
package main

import (
		"fmt"
		"bufio"
		"os"
		"strings"
		"strconv"
		"sort"
		"sync"
)


func sort_me(slice []int, wg *sync.WaitGroup){
	sort.Ints(slice)		//simple call to sort method from library
	wg.Done()						//subtracts 1 from WaitGroup counter
}


func insert(i int, ele int, slice []int) ([]int){
	//function allows us to insert an element anywhere in an 
	//existing slice
	if i > len(slice) {
		slice = append(slice, ele)
	} else {
		slice = append(slice, 0)
		copy(slice[i+1:], slice[i:])
		slice[i] = ele
	}
	return slice
}

func merge(long_slice []int, short_slice []int) ([]int) {
	//takes two slices (one if preferably longer than the other)
	//that have already been sorted individually
	//and adds the elements of the (presumeably) shorter slice one-by-one
	//to the longer slice, inserting each at the proper position to keep the
	//final product sorted
	if len(short_slice) == 0 {
		return long_slice
	} else {
		var first int
		for i := 0; i < len(long_slice) - 1; i++ {		
			if short_slice[0] >= long_slice[i] && short_slice[0] < long_slice[i + 1] {
				first, short_slice = short_slice[0], short_slice[1:]
				long_slice = insert(i + 1, first, long_slice)
				return merge(long_slice, short_slice)
			} else if short_slice[0] >= long_slice[len(long_slice) - 1]{
				first, short_slice = short_slice[0], short_slice[1:]
				long_slice = insert(len(long_slice) + 1, first, long_slice)
				return merge(long_slice, short_slice)
			} else if short_slice[0] <= long_slice[0] {
				first, short_slice = short_slice[0], short_slice[1:]
				long_slice = insert(0, first, long_slice)
				return merge(long_slice, short_slice)
			}
		}	
	}
	fmt.Println("List is not long enough!")
	return long_slice
}

func merge_all(one []int, two []int, three []int, four []int) ([]int) {
	//merges four slices and maintains sorting order
	final := merge(one, two)
	final = merge(final, three)
	final = merge(final, four)
	return final
}


func main(){
	fmt.Println("Please enter a sequence of numbers separated by spaces.")
	fmt.Println("You must enter at least eight numbers.")
	var s string
	scanner := bufio.NewScanner(os.Stdin)		//create scanner object
	if scanner.Scan() {
        s = scanner.Text()			//pull in input as strings
    }
  num_str := strings.Fields(s)
	numbers := make([]int, 0)		//this will hold our input as integers

	for _, v := range num_str {		//convert strings to integers
		a, _ := strconv.Atoi(v)
		numbers = append(numbers, a)
	}

	quotient := len(numbers) / 4		//compute min size of sub-slice
	remainder := len(numbers) % 4		//compute leftovers, they'll be added to slice 4

	one := make([]int, 0)
	two := make([]int, 0)
	three := make([]int, 0)
	four := make([]int, 0)

	/*
	Next four for loops, we load up slices 1-4

	*/

	for i := 0; i < quotient; i++ {
		var x int
		x, numbers = numbers[0], numbers[1:]
		one = append(one, x)
	}

	for i := 0; i < quotient; i++ {
		var x int
		x, numbers = numbers[0], numbers[1:]
		two = append(two, x)
	}

	for i := 0; i < quotient; i++ {
		var x int
		x, numbers = numbers[0], numbers[1:]
		three = append(three, x)
	}

	for i := 0; i < quotient + remainder; i++ {
		var x int
		x, numbers = numbers[0], numbers[1:]
		four = append(four, x)
	}

	var wg sync.WaitGroup					//we need a wait group to synchronize 
	wg.Add(4)

	/*
		One, two, three, and four can be sorted concurrently, but they sorting
		must be complete for the main function to continue...
	*/

	go sort_me(one, &wg)
	go sort_me(two, &wg)
	go sort_me(three, &wg)
	go sort_me(four, &wg)

	wg.Wait()			//end blocking of main

	
	fmt.Println("Sub-slice 1: ", one)
	fmt.Println("Sub-slice 2: ", two)
	fmt.Println("Sub-slice 3: ", three)
	fmt.Println("Sub-slice 4: ", four)

	fmt.Println("Merged slice: ", merge_all(one, two, three, four))

	
}