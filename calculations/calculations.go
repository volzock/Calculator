package calculations

import (
	"calculator/structs"
	"calculator/utils"
	"fmt"
	"strconv"
	"strings"
)

func Compute(expression string) int {
	clearExpresssion(&expression)

	for strings.Contains(expression, "(") {
		expression = fmt.Sprintf("%s%d%s",
			expression[0:strings.Index(expression, "(")],
			calculate(expression[strings.Index(expression, "(")+1:strings.Index(expression, ")")]),
			expression[strings.Index(expression, ")")+1:])
	}

	return calculate(expression)
}

func calculate(expression string) int {
	root := new(structs.BinaryTreeNode)
	processing(&expression, root)
	calculations(root)
	return root.Value
}

func clearExpresssion(expression *string) {
	*expression = strings.ReplaceAll(*expression, " ", "")
	*expression = strings.ReplaceAll(*expression, "\n", "")
}

func processing(expression *string, node *structs.BinaryTreeNode) {
	if strings.Contains(*expression, "+") || strings.Contains(*expression, "-") {
		index, expressionType := utils.Min(strings.Index(*expression, "+"),
			strings.Index(*expression, "-"),
			"+", "-")
		node.ExpressionType = expressionType

		node.Left = new(structs.BinaryTreeNode)
		leftExpression := (*expression)[0:index]
		processing(&leftExpression, node.Left)

		node.Right = new(structs.BinaryTreeNode)
		rightExpression := (*expression)[index+1:]
		processing(&rightExpression, node.Right)
	} else if strings.Contains(*expression, "/") || strings.Contains(*expression, "*") {
		index, expressionType := utils.Min(strings.Index(*expression, "*"),
			strings.Index(*expression, "/"),
			"*", "/")
		node.ExpressionType = expressionType

		node.Left = new(structs.BinaryTreeNode)
		leftExpression := (*expression)[0:index]
		processing(&leftExpression, node.Left)

		node.Right = new(structs.BinaryTreeNode)
		rightExpression := (*expression)[index+1:]
		processing(&rightExpression, node.Right)
	} else {
		node.Value, _ = strconv.Atoi(*expression)
	}
}

func calculations(node *structs.BinaryTreeNode) {
	if node.Left == nil && node.Right == nil {
		return
	} else {
		calculations(node.Left)
		calculations(node.Right)
		node.Value = doMathFunction(node.Left.Value, node.Right.Value, node.ExpressionType)
	}
}

func doMathFunction(a, b int, symbol string) int {
	switch symbol {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}
}
