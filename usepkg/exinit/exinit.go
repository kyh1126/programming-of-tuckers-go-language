package exinit

import "fmt"

var (
	a = c + b
	b = f()
	c = f()
	d = 3
)

func init() {
	d++
	fmt.Println("exinit.init function")
}

func f() int {
	d++
	fmt.Println("f() d:", d)
	return d
}

func PrintD() {
	fmt.Println("d:", d)
}
