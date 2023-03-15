// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 198.

// Package eval provides an expression evaluator.
package eval

import (
	"fmt"
	"math"
	"strconv"
)

//!+env

type Env map[Var]float64

//!-env

//!+Eval1

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

//!-Eval1

//!+Eval2

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

func (m min) Eval(env Env) float64 {
	result := m.args[0].Eval(env)
	for i := 1; i < len(m.args); i++ {
		if m.args[i].Eval(env) < result {
			result = m.args[i].Eval(env)
		}
	}
	return result
}

//!-Eval2

func (v Var) String() string { return "[" + string(v) + "]" }
func (l literal) String() string {
	return "[" + strconv.FormatFloat(float64(l), 'f', 2, 64) + "]"
}
func (u unary) String() string { return "[" + string(u.op) + u.x.String() + "]" }
func (b binary) String() string { return "[" + b.x.String() + string(b.op) + b.y.String() + "]" }
func (c call) String() string {
	result := "[" + c.fn
	for _, arg := range c.args {
		result += arg.String()
	}
	return result + "]"
}
func (m min) String() string {
	result := "[min"
	for _, arg := range m.args {
		result += arg.String()
	}
	return result + "]"
}

