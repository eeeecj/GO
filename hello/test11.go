package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var s Simpler = &Simple{}
	fmt.Printf("获得的值为:%d\n", s.get())
	s.set(3)

	var areaIntf AreaInterface
	var periIntf PeriInterface

	sq1 := new(square)
	sq1.side = 5
	tr1 := new(Triangle)
	tr1.base = 3
	tr1.height = 5

	periIntf = sq1
	fmt.Printf("The square has perimeter: %f\n", periIntf.perimeter())

	areaIntf = tr1
	fmt.Printf("The triangle has area: %f\n", areaIntf.Area())

	data := &float64Array{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	Sort(data)
	fmt.Printf("%v\n", *data)
	fmt.Println(IsSorted(data))

	data1 := NewFloat64Array()
	data1.Fill(15)
	fmt.Printf("原数组:%v\n", data1)
	Sort(&data1)
	fmt.Printf("排序后:%v\n", data1)

	ps := &Persons{Person{"kska", "aska"}, Person{"jjjj", "iiii"}}
	fmt.Printf("排序前：%v\n", *ps)
	Sort(ps)
	fmt.Printf("排序后:%v\n", *ps)

	fn := func(i obj) obj {
		switch i.(type) {
		case int:
			return i.(int) * 2
		case string:
			return i.(string) + i.(string)
		}
		return i
	}
	fmt.Printf("数字map结果为:%v\n", MapFunc(fn, []obj{0, 1, 2, 3}))
	fmt.Printf("字符串map结果为:%v\n", MapFunc(fn, []obj{"a", "b", "c"}))

	fmt.Printf("数字map结果为:%v\n", MapFuncVar(fn, 0, 1, 2, 3))
	fmt.Printf("字符串map结果为:%v\n", MapFuncVar(fn, "a", "b", "c"))

	k := new(Stack)
	k.Push(3)
	k.Push("s")
	fmt.Println(k.Top())
	fmt.Println(k.Pop())
	fmt.Println(k)
}

type Simpler interface {
	get() int
	set(int)
}
type Simple struct {
}

func (s *Simple) get() int {
	fmt.Printf("这是get\n")
	return 10
}

func (s *Simple) set(x int) {
	fmt.Printf("设置的值为：%d\n", x)
}

type Triangle struct {
	base   float64
	height float64
}

type AreaInterface interface {
	Area() float64
}

func (t *Triangle) Area() float64 {
	return t.base * t.height / 2
}

type square struct {
	side float64
}
type PeriInterface interface {
	perimeter() float64
}

func (s *square) perimeter() float64 {
	return 4 * s.side
}

type float64Array []float64

type Sorter interface {
	len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func (f *float64Array) len() int {
	return len(*f)
}

func (f *float64Array) Less(i, j int) bool {
	return (*f)[i] < (*f)[j]
}

func (f float64Array) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func Sort(data Sorter) {
	for i := 1; i < data.len(); i++ {
		for j := 0; j < data.len()-i; j++ {
			if data.Less(j+1, j) {
				data.Swap(j, j+1)
			}
		}
	}
}

func IsSorted(data Sorter) bool {
	for i := 0; i < data.len()-1; i++ {
		if data.Less(i+1, i) {
			return false
		}
	}
	return true
}

func NewFloat64Array() float64Array {
	return make(float64Array, 25)
}

func (p float64Array) List() string {
	s := "{"
	for i := 0; i < p.len(); i++ {
		if p[i] == 0 {
			continue
		}
		s += fmt.Sprintf("%3.1f,", p[i])
	}
	s += "}"
	return s
}

func (p float64Array) String() string {
	return p.List()
}

func (p float64Array) Fill(n int) {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < n; i++ {
		p[i] = (rand.Float64()) * 100
	}
}

type Person struct {
	firstName string
	lastName  string
}

type Persons []Person

func (p Persons) len() int { return len(p) }
func (p Persons) Less(i, j int) bool {
	in := p[i].firstName + " " + p[i].lastName
	jn := p[j].firstName + " " + p[j].lastName
	return in < jn
}
func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type obj interface{}

func MapFunc(fn func(obj) obj, o []obj) []obj {
	for i, v := range o {
		o[i] = fn(v)
	}
	return o
}

func MapFuncVar(fn func(obj) obj, list ...obj) []obj {
	for i, v := range list {
		list[i] = fn(v)
	}
	return list
}

type Stack []interface{}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Push(x interface{}) {
	*s = append(*s, x)
}

func (s Stack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("栈为空")
	}
	return s[s.Len()-1], nil
}
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("栈为空")
	}
	r := *s
	*s = r[:r.Len()-1]
	return r[r.Len()-1], nil
}
