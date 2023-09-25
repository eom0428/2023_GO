package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // get the current data and time as an integer
	answer := rand.Intn(100) + 1 // random number between 1 and 100
	fmt.Println("Guss Number Game")
	fmt.Println((answer))

	reader := bufio.NewReader(rd)(os.Stdin)

	fmt.Print("input guess number : ")
	inputNumber, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inputNumberString = strings.TrimSpace(inputNumberString)
	inputNumber, err := strconv.Atoi(inputNumberString)
	if err != nil {
		log.Fatal(err)
	}
	if inputNumber < answer {
		fmt.Println("Guess number is lower then answer") //Answe is higher
	}
}
