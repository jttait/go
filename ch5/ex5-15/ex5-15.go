package main

import "fmt"

func min(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}

	result := vals[0]
	for _, val := range vals {
		if result > val {
			result = val
		}
	}
	return result
}

func max (vals ...int) int {
	if len(vals) == 0 {
		return 0
	}

	result := vals[0]
	for _, val := range vals {
		if result < val {
			result = val
		}
	}
	return result
}

func min2(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("Error: Must provide at least one argument")
	}
	return min(vals...), nil
}

func max2(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("Error: Must provide at least one argument")
	}
	return max(vals...), nil
}


func main() {
	fmt.Println(min(1, 2, 3, 4))

	fmt.Println(min())

	fmt.Println(min(-10, 4, 2, -4, 16))

	fmt.Println(max(1, 2, 3, 4))

	fmt.Println(max())

	fmt.Println(max(-10, 4, 2, -4, 16))

	result, err := min2(1, 2, 3, 4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = min2()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = max2(1, 2, 3, 4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = max2()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
