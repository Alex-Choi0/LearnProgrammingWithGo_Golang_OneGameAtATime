package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Learn Programming With Go (Golang), One Game at a Time Excercise2")
	fmt.Printf("\n")
	fmt.Println("Problem 01 - Hello World")
	fmt.Println("hello World")
	fmt.Printf("\n\n")

	fmt.Println("Problem 02 - Print your name")
	fmt.Println("myName : alex")
	fmt.Printf("\n\n")

	fmt.Println("Problem 03 - Print your name and age")
	fmt.Println("myName : alex")
	fmt.Println("age : 36")
	fmt.Printf("\n\n")

	fmt.Println("Problem 04 - Print the numbers from 1 to 10")

	for i := 1; i < 11; i++ {
		fmt.Printf("%d \t", i)
	}

	fmt.Printf("\n\n")

	fmt.Println("Problem 05 - Raise 2 to the power of 11")
	var num int = 2
	var result int = 1
	for i := 1; i < 12; i++ {
		result *= num
	}
	fmt.Printf("result : %d", result)
	fmt.Printf("\n\n")
	/*
		fmt.Println("Problem 06 - Print the numbers from 1 to 1000")
		for i := 1; i < 1001; i++ {
			fmt.Printf("%d\n", i)
		}
		fmt.Printf("\n\n")
	*/

	fmt.Println("Problem 07 - Generate a random number from 0 to 10")

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Printf("%d", r1.Intn(10))
	fmt.Printf("\n\n")

	fmt.Println("Problem 08 - Print the current date")
	fmt.Println(time.Now())
	fmt.Printf("\n\n")
}
