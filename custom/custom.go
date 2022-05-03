package custom

import "fmt"

func Show(msg string, n int) {
	dots := ""

	for i := 0; i < n; i++ {
		dots += ":"
	}

	m := fmt.Sprintf("%v %v %v", dots, msg, dots)
	fmt.Println(m)
}