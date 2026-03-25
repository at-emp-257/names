package animals

import (
	"github.com/at-emp-257/names/utils"
)

var animals = []string{
	"anaconda",
	"bear",
	"cat",
}

func PrintAll() {
	utils.PrintAll("Printing animals", animals)
}