package main

import "fmt"

func main() {
	var myString string

	var hello string = "\thello\n"
	var title string = `\tHello\n`
	word := "my string" //автоматическое определение

	str := "你好"
	str = "Привет"
	fmt.Println(myString + "\n" + hello + title + word + "\n" + str)

	var b byte = 'c' // байтовое представление = 99
	var r rune = '你' // = 20320
	fmt.Println(b)
	fmt.Println(r)

	str = "12345"
	fmt.Println(str[0])   // 49
	fmt.Println(str[1])   // 50
	fmt.Println(len(str)) // 5

	str = "строка"        // 6 символов
	fmt.Println(len(str)) // 12

	//str[from:to]
	newStr := str[2:4] // "34"
	fmt.Println(newStr)
	newStr = str[:4] // "1234"
	fmt.Println(newStr)
	newStr = str[2:] // "345"
	fmt.Println(newStr)
	// from - 0, to - len(str)

	word = "Hello"
	hello = word + " world!" // "Hello world"
	//word[0] = "h" // ошибка
	s := "12345"
	s = s[:2] + "7" + s[3:]
	fmt.Println(s)

}

func varmain() {
	var v1 int //v==0
	fmt.Println(v1)

	var v2 int = 100
	fmt.Println(v2)

	v3 := 5
	fmt.Println(v3)

	var v5, v6 = 1, 2
	v5, v6 = v6, v5
	fmt.Println(v5, v6)

	var v7 int
	v5, v6, v7 = 1, 2, 3

	var (
		v01 = 1
		v02 = "string"
	)

	fmt.Println(v7)
	_ = v01
	_ = v02
	//стиль
	var EarthRadiusInMeter int = 6371000
	_ = EarthRadiusInMeter

	var earth_radius_in_meter int = 6371000
	_ = earth_radius_in_meter

	x := 111
	ptr := &x
	fmt.Println(x)
	*ptr = 55
	fmt.Println(x)

	a := 100
	p := &a

	fmt.Println(p)
}
