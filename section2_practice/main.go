package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main(){
	// Problem 01

	fmt.Println("Problem 01 - Hello World")
	fmt.Println("Hello World!")

	fmt.Printf("\n\n")

	fmt.Println("Problem 02 - Print your name")
	fmt.Println("Alex")

	fmt.Printf("\n\n")

	fmt.Println("Problem 04 - Print the numbers from 1 to 10")
	for i := 1; i < 11; i++ {

		fmt.Printf("%d\n", i);

	}


	fmt.Printf("\n\n")

	fmt.Println("Problem 05* - Raise 2 to the power of 11")
	fmt.Println(math.Pow(2, 11))

	fmt.Printf("\n\n")


	fmt.Println("Problem 06* - Print the numbers from 1 to 1000")
	for i:=1 ; i < 1000; i++ {
		fmt.Println(i)
	}
	fmt.Printf("\n\n")

	fmt.Println("Problem 07* - Generate a random number from 0 to 10")
	s1 := rand.NewSource(time.Now().UnixNano())
   r1 := rand.New(s1)
	fmt.Printf("%d\n", r1.Intn(10))

	fmt.Printf("\n\n")

	fmt.Println("Problem 08* - Print the current date")
	currentTime := time.Now()
	fmt.Println(currentTime.String())
	fmt.Printf("\n\n")

}
