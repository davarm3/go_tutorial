package main

import "fmt"

// const s string = "This is a const"

// Main function
func main() {
	// 1. Hello world!
	// fmt.Println("Hello World!")

	// 2. Values
	// fmt.Println("go" + "lang")
	// fmt.Println("1+1 = ", 1+1)
	// fmt.Println("7.0/3.0 =", 7.0/3.0)

	// fmt.Println(true && false)
	// fmt.Println(true || false)
	// fmt.Println(!true)

	// 3. Variables
	// var a = "hello"
	// fmt.Println(a)

	// var b, c = 2, 4
	// fmt.Println(b + c)

	// var d = true
	// fmt.Println(d)

	// var e int
	// fmt.Println(e)

	// f := "How are u?"
	// fmt.Println(f)

	// 4. Constants
	// fmt.Println(s)
	
	// const n = 500000000
	// const d = 3e20 / n
	// fmt.Println(d)
	// fmt.Println(int64(d))

	// 5. For-loops
	var i int = 1
	for i <= 3 {
		fmt.Println(i)
		i = i +1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 10 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// If/else
	
}
