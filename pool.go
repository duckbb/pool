package pool

import (
	"errors"
	"reflect"
	"time"
)

type Pool struct {
	objs chan interface{}
}

func NewPool(number int, option ...interface{}) (*Pool, error) {
	if number < 0 {
		return nil, errors.New("number must greater than 0")
	}
	m := make(chan interface{}, number)
	if len(option) == 0 {
		return &Pool{objs: m}, nil
	}
	if ret, ok := option[0].(func() interface{}); ok {
		for i := 0; i < number; i++ {
			m <- ret()
		}
	} else if reflect.TypeOf(option[0]).Kind() == reflect.Struct {
		temp := option[0]
		for i := 0; i < number; i++ {
			m <- temp
		}
	} else {
		return nil, errors.New("create pool fail")
	}
	return &Pool{objs: m}, nil
}

//func NewPoolFunc(number int, f func() interface{}) (*Pool, error) {
//	if number < 0 {
//		return nil, errors.New("number must greater than 0")
//	}
//	m := make(chan interface{}, number)
//	for i := 0; i < number; i++ {
//		m <- f()
//	}
//	return &Pool{objs: m}, nil
//}
//func NewPoolStruct(number int,interface{}){
//
//}

func (p *Pool) Add(obj interface{}) error {
	select {
	case p.objs <- obj:
		return nil
	default:
		return errors.New("obj add fail")
	}
}

func (p *Pool) Get(timeout time.Duration) (interface{}, error) {
	select {
	case ret, ok := <-p.objs:
		if !ok {
			return nil, errors.New("pool has closed!")
		}
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}
}

func (p *Pool) Put(obj interface{}) error {
	select {
	case p.objs <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}
