package main

import (
	"LearnGo/data_structures/05-Recursion/linkedlistr"
	"fmt"
)

func main() {
	l := linkedlistr.NewLinkedListR()
	for i := 0; i < 5; i++ {
		l.AddLast(i)
		fmt.Println(l)
	}

	l.Add(2, 666)
	fmt.Println(l)

	fmt.Println(l.GetFirst(), l.GetLast(), l.Get(2))

	l.Set(3, 222)
	fmt.Println(l)

	fmt.Println(l.RemoveFirst(), l.RemoveLast(), l.Remove(2))
	fmt.Println(l)

	fmt.Println(l.RemoveElement(66), l.RemoveElement(666))
	fmt.Println(l)
	fmt.Println(l.Contains(3), l.Contains(2), l.GetSize(), l.IsEmpty())
}
