package request

import (
	//"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Condition interface {
	IsQueryCondi() bool
	Cond2QueryObj(cate string) bson.M

	IsUpdateCondi() bool
	Cond2UpdateObj(cate string) bson.M
}
