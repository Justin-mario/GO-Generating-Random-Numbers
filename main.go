package main

import (
	rand2 "crypto/rand"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//randomNumberWithCrypto()
	//generatingDifferentRandomNumbers
	additionQuiz()
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
	correct := 0
	incorrect := 0
	answer := firstOperand + secondOperand

	fmt.Printf("what is %v + %v = ", firstOperand, secondOperand)
	fmt.Scan(&userInput)

	for i := 0; i < 11; i++ {

		if answer == userInput {
			fmt.Println("Correct!!!")
			correct++
			additionQuiz()
		} else {

			fmt.Println("incorrect")
			incorrect++
			additionQuiz()

		}

	}
}
