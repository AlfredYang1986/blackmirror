package main

import (
	_ "blackmirror/src/brand"
	"fmt"
)

func main() {
	b := brand.alfred{
		name: "alfred"}

	fmt.Println(b)
}
