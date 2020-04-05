package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calc struct {}

func (Calc) Operate(entrada string, operador string) {
	cleanedIn := strings.Split(entrada, operador)
	operationOne := parsear(cleanedIn[0])
	operationTwo := parsear(cleanedIn[1])
	switch operador {
	case "+":
		fmt.Println(operationOne + operationTwo)
	case "-":
		fmt.Println(operationOne - operationTwo)
	case "*":
		fmt.Println(operationOne * operationTwo)
	default:
		fmt.Println(operador, "No esta soportado")
	}
}

func parsear(entrada string) int {
	value, _ := strconv.Atoi(entrada)
	return value
}

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	fmt.Println("Enter your input")
	operation := ReadInput()
	fmt.Println("Enter your operation")
	operator := ReadInput()
	c := Calc{}
	c.operate(operation, operator)
}