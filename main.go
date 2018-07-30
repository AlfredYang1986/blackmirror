package main

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/brand"
)

func main() {
	a, _ := brand.GetAlfred()
	fmt.Println(a)
	fmt.Println(a.Name)
}
