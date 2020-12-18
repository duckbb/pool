# comon object pool 

## 1.install
go get -u github.com/duckbb/pool
## 2.create pool demo

````
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
````

1.default pool ,you need to add obj by yourself
```
pool, err := NewPool(10)
```
2.create pool by struct
```
pool, err := NewPool(10,&Persion{})
```
3.create pool by func
```
pool, err := NewPool(10, func() interface{} {
		rand.Seed(time.Now().UnixNano())
		t := rand.Intn(100)

		return &Persion{
			Name: "小红-" + strconv.Itoa(t),
			age:  t,
		}
	})
```


