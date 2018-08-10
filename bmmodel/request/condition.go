package request

import (
	//"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Condition interface {
	IsQueryCondi() bool
	Cond2QueryObj() bson.M

	IsUpdateCondi() bool
	Cond2UpdateObj() bson.M
}
