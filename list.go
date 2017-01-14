package main

import "fmt"

type List struct {
	Head interface{}
	Tail *List
}

func (l *List) Plus(x interface{}) List {
	return List{Head: x, Tail: l}
}

func main() {
	var lst List
	lst = lst.Plus(1).Plus(2).Plus(4)

	fmt.Println(lst.Tail.Head)
}
