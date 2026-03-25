package fruits

import (
	"github.com/at-emp-257/names/utils"
)

var fruits = []string{
	"apple",
	"banana",
	"cucumber",
}

func PrintAll() {
	utils.PrintAll("Printing fruits", fruits)
}
