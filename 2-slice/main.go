package main

import (
	"fmt"
	"math/rand"
)

func printSlice(inputSlice *[]int) {
	for i := range *inputSlice {
		fmt.Println((*inputSlice)[i])
	}
}

func sliceExample(inputSlice *[]int) *[]int {
	var returnSlice []int

	for i := range *inputSlice {
		if (*inputSlice)[i]%2 == 0 {
			returnSlice = append(returnSlice, (*inputSlice)[i])
		}
	}

	return &returnSlice
}

func addElements(inputSlice *[]int, num int) *[]int {
	*inputSlice = append(*inputSlice, num)

	return inputSlice
}

func copySlice(inputSlice *[]int) *[]int {
	returnSlice := make([]int, len(*inputSlice))
	copy(returnSlice, *inputSlice)
	return &returnSlice
}

func removeElement(inputSlice *[]int, indexRemoved int) *[]int {
	var returnSlice []int
	returnSlice = append(returnSlice, (*inputSlice)[:indexRemoved]...)
	returnSlice = append(returnSlice, (*inputSlice)[indexRemoved+1:]...)
	return &returnSlice
}

func main() {
	var originalSlice = []int{}

	fmt.Println("Сгенерированный массив:")
	for i := 0; i <= 10; i++ {
		originalSlice = append(originalSlice, rand.Intn(100))
		fmt.Println(originalSlice[i])
	}

	fmt.Println("---------------------------------")
	evenSlice := sliceExample(&originalSlice)

	fmt.Println("Чётные элементы массива:")
	printSlice(evenSlice)

	fmt.Println("---------------------------------")
	addSlice := addElements(&originalSlice, rand.Intn(100))

	fmt.Println("Массив с добавленным элементом:")
	printSlice(addSlice)

	fmt.Println("---------------------------------")
	copySliceResult := copySlice(&originalSlice)

	(*copySliceResult)[0] += 1

	fmt.Println("Исходный массив:")
	printSlice(addSlice)
	fmt.Println("\nСкопированный массив:")
	printSlice(copySliceResult)

	fmt.Println("---------------------------------")
	removedSlice := removeElement(&originalSlice, rand.Intn(len(originalSlice)))

	fmt.Println("Массив с удалённым элементом:")
	printSlice(removedSlice)
}
