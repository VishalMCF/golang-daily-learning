package event_examples

import (
	"fmt"
	"github.com/agoalofalife/event"
	"strconv"
)

type Hero struct {
	weapon string
	age    int
}

func TestStructAsKey() {
	hero := Hero{
		"sword",
		10923,
	}

	e := event.New()

	e.Add(hero, func(hero Hero) {
		fmt.Println("hero weapon is " + hero.weapon + " and he is " + strconv.Itoa(hero.age) + " years old")
	}, hero)

	fmt.Println(e.Fire(hero))
}
