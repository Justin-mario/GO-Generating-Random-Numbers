package main

import (
	rand2 "crypto/rand"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var correct int
var incorrect int
var count = 1

func main() {
	//randomNumberWithCrypto()
	//generatingDifferentRandomNumbers
	additionQuiz()
	multiplicationQuiz()
}

func generatingDifferentRandomNumbers() {
	x1 := rand.Int()
	x2 := rand.Intn(200)
	x3 := (rand.Float64() * 7) + 8
	x4 := rand.NewSource(time.Now().UnixNano())
	x5 := rand.New(x4)

	fmt.Printf("random is %v\n%v\n%v\n%v\n", x1, x2, x3, x5)
	fmt.Printf("with time %v\n", x4)
}

func randomNumberWithCrypto() {
	RandomCrypto, _ := rand2.Prime(rand2.Reader, 10)
	x1 := rand.Int()
	x2 := rand.Intn(19)
	x6 := rand.Perm(x2)
	fmt.Printf("with crypto %d\n%v%d\n%v", RandomCrypto, x6, x2, x1)
}

func additionQuiz() {
	firstOperand := rand.Intn(19)
	secondOperand := rand.Intn(18)
	var userInput int

	answer := firstOperand + secondOperand

	fmt.Printf("what is %v + %v = ", firstOperand, secondOperand)
	_, err := fmt.Scan(&userInput)
	if err != nil {
		return
	}
	if count == 10 {
		fmt.Printf("Thank you for Taking the Quiz Your got %v right and failed %v ", correct, incorrect)
		os.Exit(0)
	}
	if userInput == -1 {
		fmt.Printf("Thank you for Taking the Quiz Your got %v right and failed %v ", correct, incorrect)
		os.Exit(0)

	} else if answer == userInput {
		fmt.Println("Correct!!!")
		correct = correct + 1
		fmt.Println(count)
	} else {
		fmt.Println("incorrect")
		incorrect++
	}
	count++
	additionQuiz()

}

func multiplicationQuiz() {
	firstOperand := rand.Intn(19)
	secondOperand := rand.Intn(18)
	var userInput int

	answer := firstOperand * secondOperand

	fmt.Printf("what is %v * %v = ", firstOperand, secondOperand)
	_, err := fmt.Scan(&userInput)
	if err != nil {
		return
	}
	if count == 10 {
		fmt.Printf("Thank you for Taking the Quiz Your got %v right and failed %v ", correct, incorrect)
		os.Exit(0)
	}
	if userInput == -1 {
		fmt.Printf("Thank you for Taking the Quiz Your got %v right and failed %v ", correct, incorrect)
		os.Exit(0)

	} else if answer == userInput {
		fmt.Println("Correct!!!")
		correct = correct + 1
		fmt.Println(count)
	} else {
		fmt.Println("incorrect")
		incorrect++
	}
	count++
	additionQuiz()

}
