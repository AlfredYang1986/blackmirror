package order

import (
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"fmt"
)

func (bd *Order) FindMulti(req request.Request) (interface{}, error) {
	var rs []Order
	err := bmmodel.FindMutil(req, &rs)
	var mrs []Order
	for _,r := range rs {
		r.ResetIdWithId_()
		mrs = append(mrs, r)
	}
	fmt.Println(mrs)
	return mrs, err
}