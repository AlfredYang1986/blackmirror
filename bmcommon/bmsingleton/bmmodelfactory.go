package bmsingleton

import (
	"errors"
	"reflect"
	"sync"
)

type modelFactory struct {
	bm map[string]reflect.Type
}

var factory *modelFactory
var once sync.Once

func GetFactoryInstance() *modelFactory {
	once.Do(func() {
		factory = &modelFactory{
			bm: make(map[string]reflect.Type)}
	})

	return factory
}

func (f *modelFactory) RegisterModel(name string, tp interface{}) {
	for k, _ := range f.bm {
		if k == name {
			return
		}
	}

	t := reflect.TypeOf(tp).Elem()
	f.bm[name] = t
}

func (f *modelFactory) ReflectInstance(name string) (interface{}, error) {
	var tp reflect.Type
	b := false
	for k, v := range f.bm {
		if k == name {
			tp = v
			b = true
		}
	}

	if b == true {
		reval := reflect.New(tp).Elem().Interface()
		return reval, nil
	} else {
		return 0, errors.New("not register class")
	}
}
