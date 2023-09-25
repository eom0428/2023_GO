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
	inputNumberString = strings.TrimSpace(inputNumberString) // remove space bar
	inputNumber, err := strconv.ParseFloat(inputNumberString, 32)

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
