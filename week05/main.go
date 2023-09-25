package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("input Score : ")
	reader := bufio.NewReader(os.Stdin)
	inputNumber, err := reader.ReadString('\n') // option 2
	if err != nil {
		log.Fatal(err)
	}
	input = string.TrimSpace(input)
	inputNumberString = strings.TrimSpace(inputNumberString)
	inputNumber, err := strconv.ParseFloat(s, bitSize)(inputNumberString, 32)

	if err != nil {
		log.Fatal(err)
	}
	var grade string
	if inputNumber >= 90 {
		grade = "A grade!"
	} else {
		grade = "BCDE grade~"
	}
	fmt.Println(" you got" + grade)
}
