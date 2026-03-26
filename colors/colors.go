package colors

import (
	"github.com/at-emp-257/names/utils"
)

var colors = []string{
	"aqua",
	"black",
	"cyan",
}

func List() {
	utils.PrintAll("Printing colors", colors)
}
