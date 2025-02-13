package main

import (
	"fmt"

	"github.com/hartmutobendorf/go_mod_deps/go_mod/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
