package main

import (
	"fmt"
	"math"
	"strings"
)

var history string = ""

type Calculator struct {
	a  float64
	b  float64
	op string
}

func main() {
	prompts()
}

func prompts() {
	var a, b, res float64
	var sgn string
	fmt.Println("Enter first value: ")
	fmt.Scanln(&a)
	fmt.Println("Enter second value: ")
	fmt.Scanln(&b)
	fmt.Println("Choose Operator: add:+, Sub:-, div:/, multi:*, modu:%  ||| To end press X")
	fmt.Scanln(&sgn)
	if strings.ToLower(sgn) != "x" {
		//make calculator and calculate
		// return result of calculate
		//while sgn or a or b is not equals to x
		// always return result, ask new val and sign
		clc := Calculator{
			a:  a,
			b:  b,
			op: sgn,
		}
		res = calculate(clc)
		fmt.Println("Result is: ", res)
		for strings.ToLower(sgn) != "x" && strings.ToLower(sgn) != "a" {
			fmt.Println("Choose Operator: add:+, Sub:-, div:/, multi:*, modu:% ||| To end press X / Start Over press A")
			fmt.Scanln(&sgn)
			if strings.ToLower(sgn) != "x" && strings.ToLower(sgn) != "a" {
				fmt.Println("Enter second value")
				fmt.Scanln(&b)
				clc = Calculator{
					a:  res,
					b:  b,
					op: sgn,
				}
				res = calculate(clc)
				fmt.Println("Result is: ", res)
			} else {
				if sgn == "a" {
					fmt.Println("Calculator history= ", history)
					history = ""
					prompts()
				} else if sgn == "x" {
					fmt.Println("Calculator history= ", history)
					history = ""

				} else {
					end()
				}
			}

		}
	}

}

func calculate(cal Calculator) float64 {
	//c:=""
	var oper float64
	//s:=1
	//for s>=1 {
	// fmt.Println("Enter 1 for Addition: ")
	// fmt.Println("Enter 2 for Multiplication: ")
	// fmt.Println("Enter 3 for Division: ")
	// fmt.Println("Enter 4 for Subtraction: ")
	// fmt.Print("Enter 5 for Exit: ")
	// fmt.Scanf("%d", &c)
	switch cal.op {
	case "+":
		oper = cal.a + cal.b
		fmt.Println(oper)
		history += fmt.Sprintf("%f", cal.a) + " " + cal.op + " " + fmt.Sprintf("%f", cal.b) + ", "
	case "*":
		oper = cal.a * cal.b
		fmt.Println(oper)
		history += fmt.Sprintf("%f", cal.a) + " " + cal.op + " " + fmt.Sprintf("%f", cal.b) + ", "
	case "/":
		oper = cal.a / cal.b
		fmt.Println(oper)
		history += fmt.Sprintf("%f", cal.a) + " " + cal.op + " " + fmt.Sprintf("%f", cal.b) + ", "
	case "-":
		oper = cal.a - cal.b
		fmt.Println(oper)
		history += fmt.Sprintf("%f", cal.a) + " " + cal.op + " " + fmt.Sprintf("%f", cal.b) + ", "
	case "%":
		oper = math.Mod(cal.a, cal.b)
		fmt.Println(oper)
		history += fmt.Sprintf("%f", cal.a) + " " + cal.op + " " + fmt.Sprintf("%f", cal.b) + ", "
		//    case "x":
		// 	  	s = 0
		// 	  	break
	default:
		fmt.Println("Enter valid operator.")
		prompts()
	}
	// }

	return oper

}

func end() {
	fmt.Println("See you later")
}
