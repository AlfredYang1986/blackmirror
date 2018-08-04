package main

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/jsonapi"
)

var req string = `{
	"data":
		{	"id": "request id",
			"type":"request",
			"attributes": {
				"res": "Brand"
			},
			"relationships": {
				"conditions": {
					"data": [
					{
						"id": "condi 01",
						"type": "eq_cond"
					},
					{
						"id": "condi 02",
						"type": "eq_cond"
					}
					]
				}
			}
		},
		"included":[
		{
			"id": "condi 01",
			"type": "eq_cond",
			"attributes": {
				"key": "name",
				"val": "alfredyang"
			}
		},
		{
			"id": "condi 02",
			"type": "eq_cond",
			"attributes": {
				"key": "about",
				"val": "about brand"
			}
		}
		]
	}`

func main() {
	rst, _ := jsonapi.FromJsonAPI(req)
	t := rst.(request.Request)
	for _, itm := range t.Cond {
		fmt.Println(itm)
		eq := itm.(request.EQCond)
		fmt.Println(eq.Id)
		fmt.Println(eq.Ky)
		fmt.Println(eq.Vy)
	}
	fmt.Println(rst)
}
