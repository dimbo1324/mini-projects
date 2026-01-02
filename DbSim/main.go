package main

import (
	"github.com/dim4d/DbSim/terminal"
)

// var tb *typebox.TypeBox = typebox.NewTypeBox()

var ter terminal.Terminal = terminal.Terminal{}

func main() {

	str := "SET b FLOAT 3.5"

	ter.Parser(str)

}
