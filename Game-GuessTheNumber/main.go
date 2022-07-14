// This is a game where user has to guess the right number between 0 to 15.
// The user will get 5 chance to guess the number.
// As an input the user will pass the 5 values at a same time.
// Each time the function triggers, it will generate the guess number and allow the user to guess
// the number in 5 attempt.
// If user can guess it in 1st attempt, then he will be considered as very lucky.
// If user failed to guess the number, he will be categorized as unlucky.
// The package used in this code is:
// 		*math/rand support the functions which can generate the random number. https://pkg.go.dev/math/rand
//		*time get the time and return the unix time that will be input to Seed function. https://pkg.go.dev/time
//
// Seed will counter the consequences of pseudo random number generator.
// Here, UnixNano is used as an input to Seed function to generate the random number.

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	arg := os.Args
	if len(arg) != 6 {
		fmt.Println("You have to provide exactly 5 inputs as your guess.")
		return
	}
	t := time.Now()
	rand.Seed(t.UnixNano())

	guess := rand.Intn(16)
	correct := false
	for i, num := range arg {
		if i == 0 {
			i++
			continue
		}

		if val, err := strconv.Atoi(num); err != nil {
			fmt.Printf("%s", err)
		} else {
			if val > 15 {
				fmt.Println("Your choice is not between 0 and 15")
				fmt.Printf("Skipping your %d attempt\n", i)
				i++
				continue
			}
			if val == guess {
				if i == 1 {
					fmt.Println("Congrats you Guessed correctly !!!")
					fmt.Println("You are eligible for special award.")
				} else {
					fmt.Println("Congrats you Guessed correctly !!!")
				}
				fmt.Printf("You have guessed the number in %d attempt\n", i)
				correct = true
				break
			}
		}
		i++
	}
	if !correct {
		fmt.Println("You exceed your maximum attempt.")
		fmt.Println("You didnot guess the number.")
	}
	fmt.Printf("The guess number is %d\n", guess)
}
