package base

import "fmt"

func Example() {
	str := "Hello小风"
	fmt.Println(len([]rune(str)))
	for i, v := range []rune(str) {
		fmt.Println(i, string(v))
	}
}
