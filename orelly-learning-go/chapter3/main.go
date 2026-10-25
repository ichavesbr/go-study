package chapter3

import "fmt"

func Main() {
	fmt.Println("=================")
	fmt.Println("MODULE: chapter 3\n")

	var x = []int{10, 20, 30}
	fmt.Println("x:", x)
	fmt.Println("Slice no index 1:", x[1])
	x = append(x, 25, 26, 27)
	fmt.Println("novo x:", x, "\n")

	var y []int
	fmt.Println("\ny:", y)
	fmt.Println("y é nil:", y == nil)
	y = append(y, 25)
	fmt.Println("novo y:", y)
	y = append(y, x...)
	fmt.Println("outro novo y:", y, "\n")

	// capacidade
	var a []int
	fmt.Println(a, len(a), cap(a))
	a = append(a, 10)
	fmt.Println(a, len(a), cap(a))
	a = append(a, 20)
	fmt.Println(a, len(a), cap(a))
	a = append(a, 30)
	fmt.Println(a, len(a), cap(a))
	a = append(a, 40)
	fmt.Println(a, len(a), cap(a))
	a = append(a, 50)
	fmt.Println(a, len(a), cap(a), "\n")

	// criar slice com capacidade pré-definida
	b := make([]int, 5, 10)
	fmt.Println(b, len(b), cap(b), "\n")

	// limpar slice
	c := []string{"um", "dois", "três"}
	fmt.Println(c, len(c), cap(c))
	clear(c)
	fmt.Println(c, len(c), cap(c), "\n")

	// copiar slice
	d := []int{1, 2, 3, 4}
	fmt.Println("d:", d, len(d), cap(d))
	e := make([]int, 4)
	fmt.Println("e:", e, len(e), cap(e))
	num := copy(e, d)         // copia para "e" o conteúdo de "d"
	fmt.Println(e, num, "\n") // copy retorna o número de elementos copiados
}
