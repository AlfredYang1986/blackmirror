package main

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmmodel/brand"
)

var sjson string = `{"_id": "id", "name": "alfredyang",	"slogan": "i am slogan", "highlights": ["abc", "456", "789"], "about": "about brand", "awards": {"a": "1"},	"attends": {"a": "1"}, "qualifier": {"a": "1"}}`

func main() {
	a, _ := brand.FromJson(sjson)
	fmt.Println(a.GetName())
	fmt.Println(a.GetHighlights())
}
