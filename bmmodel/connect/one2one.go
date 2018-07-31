package connect

import ()

type one2one interface {
	queryConnect(a interface{}, k string) (interface{}, error)
}
