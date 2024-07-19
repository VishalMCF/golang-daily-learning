package main

import (
	"fmt"
)

type Student struct {
	a, b, c int
}

func main() {
	// event_examples.OneToMany()
	// event-examples.ExampleBaseServer()
	// event-examples.ExampleEventAndChannels()
	// mutexes_concept.Demo()
	// atomic_concept.NonAtomicDemo()
	// atomic_concept.WithAtomicDemo()
	// reflection.Print()
	// reflection.TestReflection()
	// event_examples.TestStructAsKey()
	s := Student{
		a: 3,
		b: 5,
		c: 1,
	}
	a := make(map[string]int)
	a["0"] = 1
	fmt.Println("Hello World ", a["0"], s)
}
