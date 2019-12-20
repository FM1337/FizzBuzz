package main

import "fmt"

func main() {
	Fizzes := []int{3}
	Buzzes := []int{5}
	for i := 1; i <= 100; i++ {
		output := ""
		decider(&Fizzes, &i, &output, "Fizz")
		decider(&Buzzes, &i, &output, "Buzz")
		if output == "" {
			output = fmt.Sprintf("%d", i)
		}
		fmt.Println(output)
	}
}

func decider(list *[]int, i *int, output *string, word string) {
	for _, n := range *list {
		if *i%n == 0 {
			*output += word
		}
	}
}
