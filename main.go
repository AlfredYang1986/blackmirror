package main

import (
	//"encoding/json"
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmmodel/brand"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	//"os"
)

var sjson string = `{
	"data":[
		{	"id": "i am id",
			"type":"brand",
			"attributes": {
				"name": "alfredyang",
				"slogan": "i am slogan",
				"about": "about brand",
				"highlights": ["abc", "456", "789"],
				"awards": {"a": "1"},
				"attends": {"a": "1"},
				"qualifier": {"a": "1"}
			},
			"relationships": {
				"locations": {
					"data": [
					{
						"id": "loc id 01",
						"type": "location"
					},
					{
						"id": "loc id 02",
						"type": "location"
					}
					]
				}
			}
		},
		{	"id": "i am id 999",
			"type":"brand",
			"attributes": {
				"name": "liuying",
				"slogan": "i am slogan",
				"about": "about brand",
				"highlights": ["abc", "456", "789"],
				"awards": {"a": "1"},
				"attends": {"a": "1"},
				"qualifier": {"a": "1"}
			},
			"relationships": {
				"location": {
					"data":
					{
						"id": "loc id 03",
						"type": "location"
					}
				}
			}
		}
		],
		"included":[
		{
			"id": "test id 01",
			"type": "test",
			"attributes": {
				"title": "test title"
			}
		},
		{
			"id": "loc id 01",
			"type": "location",
			"attributes": {
				"title": "loc title",
				"address": "beijingshi, chinese",
				"district": "fuck",
				"a": "1"
			},
			"relationships": {
				"test": {
					"data":
					{
						"id": "test id 01",
						"type": "test"
					}
				}

			}
		},
		{
			"id": "loc id 03",
			"type": "location",
			"attributes": {
				"title": "loc title",
				"address": "beijingshi, chinese",
				"district": "fuck",
				"a": "1"
			},
			"relationships": {
				"test": {
					"data":
					{
						"id": "test id 01",
						"type": "test"
					}
				}

			}
		},
		{
			"id": "loc id 02",
			"type": "location",
			"attributes": {
				"title": "loc title",
				"address": "beijingshi, chinese",
				"district": "fuck",
				"a": "1"
			}
		}
		]
	}`

func main() {
	rst, _ := jsonapi.FromJsonAPI(sjson)
	fmt.Println(rst)

	t := rst.([]interface{})
	tmp := t[0].(brand.Brand)
	reval, _ := jsonapi.ToJsonAPI(&tmp)
	fmt.Println(reval)

	//enc := json.NewEncoder(os.Stdout)
	//enc.Encode(tmp)
}
