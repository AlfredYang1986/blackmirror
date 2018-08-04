package request

import (
	//"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Condition interface {
	Cond2QueryObj() bson.M
}
