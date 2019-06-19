package main

import (
	"github.com/tbmale/iup"
)

func main() {
	iup.Open()
	defer iup.Close()

	iup.Message("IupMessage Example", "Press the button")
}
