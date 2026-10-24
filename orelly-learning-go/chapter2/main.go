package chapter2

import "fmt"

const x int64 = 10
const z = 20 * 10

const (
	idKey   = "id"
	nameKey = "name"
)

func Main() {
	fmt.Println("=================")
	fmt.Println("MODULE: chapter 2")
	const y = "Hello"

	fmt.Println(x, y, z)
	fmt.Println(idKey, nameKey)
}
