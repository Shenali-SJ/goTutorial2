package main

import "fmt"

func main() {

	var a [3]int  //declaration
	a[0] = 33  //initializing
	a[1] = 12
	a[2] = 87
	fmt.Println(a)

	//shorthand declaration

	friends := [3]string{"Rachel", "Monica", "Phoebe"}
	fmt.Println(friends)

	//only the first element is initialized during declaration
	a1 := [3]int{32}
	fmt.Println(a1)

	//size is not specified
	friends2 := [...]string{"Chandler", "Ross", "Joey"}
	fmt.Println(friends2)
	fmt.Println("Size of the friends2 array is ", len(friends2))

	//pass by value
	a2 := [...]int{34, 355, 80, 10, 28}
	fmt.Println("Before: ", a2)
	modifyArray(a2)
	fmt.Println("After: ", a2)

	var a3 = [4]float64{33.3, 10.3, 52.53, 342.3}
	iterate(a3)

	sum := iterateRange(a3)
	fmt.Println("Sum of all the elements: ", sum)

	//using blank identifier
	fruits := [4]string{"Apple", "Orange", "Kiwi"}
	for index, value := range fruits {
		fmt.Printf("Index - %d, value - %s \n", index, value)
	}

	//2d arrays
	var colours [3][2]string
	colours[0][0] = "Black"
	colours[0][1] = "White"
	colours[1][0] = "Red"
	colours[1][1] = "Blue"
	colours[2][0] = "Yellow"
	colours[2][1] = "Green"

	printArray(colours)

	fmt.Println()

	//2d arrays - shorthand declaration
	countries := [3][2]string{
		{"USA", "Canada"},
		{"India", "Sri Lanka"},
		{"Norway", "Sweden"},   //last comma is needed
	}
	printArray(countries)

	fmt.Println()

	//slices
	var slice []int = a2[2:4]
	fmt.Println(slice)

	// array with slice reference
	slice2 := []int {23, 41, 60, 39}
	fmt.Println(slice2)

	//modifying slice will modify the original array
	slice3 := changeArray([7]int{10, 20, 30, 40, 50, 60, 70})
	fmt.Printf("Length - %d, Capacity - %d \n", len(slice3), cap(slice3))

	//re-slicing
	slice3 = slice3[:cap(slice3)]
	fmt.Printf("After re-slicing... Length - %d, Capacity - %d \n", len(slice3), cap(slice3))

	//make
	slice4 := make([]int, 3, 5)
	fmt.Println(slice4)

	slice4 = append(slice4, 55)
	fmt.Println(slice4)

	//append

	//appending elements to a nil slice
	var words []string
	if words == nil {
		fmt.Println("Nil slice")
		words = append(words, "Random", "Sky")
		fmt.Println("After appending, ", words)
	}

	//appending a slice to a slice
	goodFood := []string{"Pizza", "Spaghetti"}
	badFood := []string{"Papaya", "String hoppers"}
	food := append(goodFood, badFood...)
	fmt.Println(food)

	//multidimensional slices
	slice5 := [][]int{
		{2, 45},
		{34},
		{200, 78, 90},
	}
	fmt.Println(slice5)

	//memory management

	//1. create a slice with an array reference
	animals := []string{"Cat", "Dog", "Monkey", "Ostrich", "Robin", "Parrot"}
	//2. slicing the array
	birds := animals[3:]
	birdsCopy := make([]string, len(birds))
	//3 copying
	copy(birds, birdsCopy)
}

func modifyArray(array [5]int) {
	array[0] = 0
	fmt.Println("Inside the function, ", array)
}

func iterate(array [4]float64) {
	for i := 0; i < len(array); i++ {
		fmt.Printf("array[%d] = %f \n",i , array[i])
	}
}

func iterateRange(array [4]float64) (sum float64) {
	sum = float64(0)
	for index, value := range array {
		fmt.Printf("Element %d in array is %f \n", index, value)
		sum += value
	}
	return
}

func printArray(array [3][2]string) {
	for _, v1 := range array {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func changeArray(array [7]int) (slice []int) {
	slice = array[2:5]
	fmt.Println("Before ", array)
	for i := range slice {
		slice[i]++
	}
	fmt.Println("After ", array)
	return
}
