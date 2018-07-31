package main

import (
	"fmt"
	//"github.com/alfredyang1986/blackmirror/bmmodel/brand"
	//"github.com/alfredyang1986/blackmirror/adt"
	"github.com/alfredyang1986/blackmirror/jsonapi"
)

/*var sjson string = `{"data":*/
//{"id": "i am id",
//"type":"brand",
//"attributes": {
//"name": "alfredyang",
//"slogan": "i am slogan",
//"about": "about brand",
//"highlights": ["abc", "456", "789"],
//"awards": {"a": "1"},
//"attends": {"a": "1"},
//"qualifier": {"a": "1"}
//}
//}}`

var sjson string = `{"data":
[
{"id": "i am id",
	"type":"brand",
	"attributes": {
		"name": "alfredyang",
		"slogan": "i am slogan",
		"about": "about brand",
		"awards": {"a": "1"},
		"attends": {"a": "1"},
		"qualifier": {"a": "1"}
	}
},
{"id": "i am id 22222",
	"type":"brand",
	"attributes": {
		"name": "yangyuan",
		"slogan": "i am slogan",
		"about": "about brand",
		"awards": {"a": "1"},
		"attends": {"a": "1"},
		"qualifier": {"a": "1"}
	}
}
]}`

func main() {
	rst, _ := jsonapi.FromJsonAPI(sjson)
	fmt.Println(rst)
}
