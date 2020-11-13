package pool

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

type Persion struct {
	Name string
	age  int
}

func GetPersion() *Persion {
	rand.Seed(time.Now().UnixNano())
	t := rand.Intn(100)

	return &Persion{
		Name: "小红-" + strconv.Itoa(t),
		age:  t,
	}
}

// time out
func TestNumberPool(t *testing.T) {
	pool, err := NewPool(10, GetPersion())
	if err != nil {
		t.Fatal("pool create fail", err)
	}

	for i := 0; i < 11; i++ {
		ret, err := pool.Get(time.Second)
		if err != nil {
			t.Fatal("get obj fail :", err)
		}
		if p1, ok := ret.(*Persion); ok {
			t.Log("姓名:", p1.Name, "年龄:", p1.age)
		} else {
			t.Fatal("change type fail")
		}
	}
}

//func to CreateObj
func TestFuncPool(t *testing.T) {
	pool, err := NewPool(10, func() interface{} {
		rand.Seed(time.Now().UnixNano())
		t := rand.Intn(100)

		return &Persion{
			Name: "小红-" + strconv.Itoa(t),
			age:  t,
		}
	})
	if err != nil {
		t.Fatal("pool create fail", err)
	}
	for i := 0; i < 10; i++ {
		ret, err := pool.Get(time.Second)
		if err != nil {
			t.Fatal("get obj fail :", err)
		}
		if p1, ok := ret.(*Persion); ok {
			t.Log("姓名:", p1.Name, "年龄:", p1.age)
		} else {
			t.Fatal("change type fail")
		}
	}
}

//normal obj add
func TestNewPool(t *testing.T) {
	pool, err := NewPool(10)
	if err != nil {
		t.Fatal("pool create fail")
	}
	p2 := GetPersion()
	t.Log(p2)
	err = pool.Add(p2)
	if err != nil {
		t.Fatal("add fail:", err)
	}
	b, err := pool.Get(time.Second)
	if err != nil {
		t.Fatal("get obj fail:", err)
	}

	if p1, ok := b.(*Persion); ok {
		t.Log("姓名:", p1.Name, "年龄:", p1.age)
	} else {
		t.Fatal("change type fail")
	}

}

//put obj
func TestPutObj(t *testing.T) {
	pool, err := NewPool(10, func() interface{} {
		rand.Seed(time.Now().UnixNano())
		t := rand.Intn(100)

		return &Persion{
			Name: "小红-" + strconv.Itoa(t),
			age:  t,
		}
	})
	if err != nil {
		t.Fatal("pool create fail", err)
	}
	var objs []interface{}
	for i := 0; i < 10; i++ {
		obj, err := pool.Get(time.Second)
		if err != nil {
			t.Fatal("get obj fail")
		}
		objs = append(objs, obj)
		fmt.Println("i=", i)
	}
	t.Log("objs length:", len(objs))
	//put
	for index, v := range objs {
		fmt.Println("index=", index)
		err := pool.Put(v)
		if err != nil {
			t.Fatal("put obj fail", err)
		}
	}
}
