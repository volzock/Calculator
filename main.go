package main

import (
	"bufio"
	"calculator/calculations"
	"fmt"
	"os"
)

func getExpression() string {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	expression, _ := reader.ReadString('\n')
	return expression
}

func main() {
	expression := getExpression()
	result := calculations.Compute(expression)
	fmt.Println(result)
}
