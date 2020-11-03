package main

import (
	"fmt"
	"github.com/SomtochiAma/ctci-algorithms/data-structures/stringbuilder"
)

func main() {
	s := stringbuilder.NewSB()
	s.Append("me and ")
	s.Append("you")
	fmt.Println(s.String())
}
