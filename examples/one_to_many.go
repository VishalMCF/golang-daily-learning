package examples

import (
	"fmt"
	"github.com/agoalofalife/event"
)

func OneToMany() {

	e := event.New()

	e.Add("receiving.message", func() {
		fmt.Println("Post email")
	})

	e.Add("receiving.message", func() {
		fmt.Println("Post in chat")
	})
	e.Add("receiving.message", func() {
		fmt.Println("Create task")
	})

	e.Go("receiving.message")
}
