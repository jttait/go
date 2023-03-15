// Exercise 7.15: Write a program that reads a single expression from the
// standard input, prompts the user to provide values for any variables, then
// evaluates the expression in the resulting environment. Handle all errors
// gracefully.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jttait/gopl.io/ch7/eval"
)

func main() {
	input := os.Args[1]
	expr, env, err := parseAndCheck(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(env)
	fmt.Printf("Result = %v\n", expr.Eval(env))
}

func parseAndCheck(s string) (eval.Expr, eval.Env, error) {
	if s == "" {
		return nil, nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, nil, err
	}

	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, nil, err
	}

	env := eval.Env{}
	env = make(map[eval.Var]float64)
	for v := range vars {
		fmt.Printf("Enter value for %s: ", v)
		var text string
		fmt.Scanf("%s", &text)
		f, err := strconv.ParseFloat(text, 64)
		if err != nil {
			fmt.Println(err)
		}
		env[v] = f
	}

	return expr, env, nil
}
