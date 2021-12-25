package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter first digit: ")
	input1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value must be a number")
	}

	fmt.Print("Enter second digit: ")
	input2, _ := reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value must be a number")
	}

	sum := float1 + float2
	rounded := math.Round(sum*100) / 100

	fmt.Printf("Sum of %v and %v is %v\n", float1, float2, sum)
	fmt.Printf("Rounded number is %v\n", rounded)

}
