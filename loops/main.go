package main

import "fmt"

func main() {

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println("For loop iteration:", i)
	}
	fmt.Println()
	// while loop
	i := 0
	for i < 5 {
		fmt.Println("While loop iteration:", i)
		i++
	}
	fmt.Println()
	// range loop
	slice := []string{"apple", "banana", "cherry"}
	for index, value := range slice {
		fmt.Printf("Range loop iteration %d: %s\n", index, value)
	}
	fmt.Println()
}
