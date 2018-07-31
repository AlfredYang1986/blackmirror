package adt

import (
	"errors"
)

type Stack struct {
	elems []interface{}
	count int
}

func StackInstance() Stack {
	return Stack{}
}

func (sk *Stack) PushElement(a interface{}) {
	if sk.count+1 < cap(sk.elems) {
		sk.elems[sk.count] = a
	} else {
		sk.elems = append(sk.elems, a)
	}
	sk.count++
}

func (sk *Stack) PopElement() (interface{}, error) {

	var err error
	if len(sk.elems) == 0 {
		err = errors.New("try to pop an empty stack")
		panic(err)
		return -1, nil
	}

	rst := sk.elems[sk.count-1]
	sk.count--

	return rst, nil
}

func (sk *Stack) Length() int {
	return sk.count
}

func (sk *Stack) ElemAtIndex(i int) interface{} {
	return sk.elems[i]
}
