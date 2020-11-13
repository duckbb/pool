package pool

import (
	"testing"
)

func TestNewPool(t *testing.T) {
	var a interface{} = 10
	var b int
	a = 20
	t.Log("a=", a)
	t.Log("b=", b)
}
