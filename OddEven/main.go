package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	s := ""
	numbers := []int{}

	fmt.Print("Enter size of slice: ")
	fmt.Scanln(&s)
	size, error := strconv.Atoi(s)

	if error != nil {
		fmt.Println(error)

	} else {
		for i := 0; i < size; i++ {
			numbers = append(numbers, returnRandom(11))
		}

		for _, n := range numbers {
			fmt.Print(n, " : ")

			if n%2 == 0 {
				fmt.Println("Even")
			} else {
				fmt.Println("Odd")
			}
		}
	}

}

func returnRandom(i int) int {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	r := rand.New(source)

	return (r.Intn(i))
}
