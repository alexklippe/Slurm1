package main

import "fmt"

func main() {

	var x1 [5]int

	var arr [3]int = [3]int{1, 2, 3}
	var arr2 = [3]int{4, 5, 6}
	arr3 := [3]int{7, 8, 9}
	arr5 := [...]int32{45, 46, 47, 48, 49}
	s := len(arr5)
	fmt.Println(x1, arr, arr2, arr3, arr5, s)
}
