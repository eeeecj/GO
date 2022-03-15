package main

import (
	"fmt"
	"math"
	"strings"
)

type Person struct {
	firstName string
	lastname  string
}

type Point struct {
	x, y float64
}

type Retangle struct {
	width  float64
	height float64
}

func upPerson(p Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastname = strings.ToUpper(p.lastname)
}

func Abs(p Point) float64 {
	return math.Sqrt(float64(p.x*p.x + p.y*p.y))
}

func Scale(p Point, f float64) (q Point) {
	q.x = p.x * f
	q.y = p.y * f
	return
}

func Area(R *Retangle) float64 {
	return R.height * R.width
}

func Perimeter(R *Retangle) float64 {
	return 2 * (R.width + R.height)
}

func main() {
	var person Person
	person.firstName = "sss"
	person.lastname = "aaaa"
	upPerson(person)
	fmt.Printf("firstname:%s;lastname:%s\n", person.firstName, person.lastname)

	person2 := new(Person)
	person2.firstName = "asjdh"
	person2.lastname = "skak"
	fmt.Printf("firstname:%s;lastname:%s\n", person2.firstName, person2.lastname)

	p1 := new(Point)
	p1.x = 2
	p1.y = 5
	fmt.Printf("距离为：%f;scale为X:%f,Y:%f\n", Abs(*p1), (Scale(*p1, 0.5)).x, Scale(*p1, 0.5).y)

	r1 := new(Retangle)
	r1.height = 4
	r1.width = 6
	fmt.Printf("面积为：%f;周长为:%f\n", Area(r1), Perimeter(r1))

	fmt.Printf("求和为:%d\n", vse{1, 2, 3, 4}.Sum())

	fmt.Printf("增涨工资为:%f\n", (&employee{10}).RaiseSalary(0.2))

	m := new(Mercedes)
	m.car.engine = nil
	m.car.wheelCount = 4
	m.sayHiToMerkel()
	m.numberOfWheels()
	m.GoToWorkIn()

	temp := (&Employee2{Person2{Base{2}, "aaa", "sss"}, 4}).Id()
	fmt.Printf("职员工资为：%d\n", temp)

	st1 := new(Stack)
	fmt.Printf("%v\n", st1.stack)
	st1.Push(0)
	fmt.Printf("%v\n", st1.stack)
	st1.Push(7)
	fmt.Printf("%v\n", st1.stack)
	st1.Push(10)
	fmt.Printf("%v\n", st1.stack)
	st1.Push(99)
	fmt.Printf("%v\n", st1.stack)
	p := st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
}

type vse []int

func (v vse) Sum() (s int) {
	for _, v := range v {
		s += v
	}
	return s
}

type employee struct {
	salary float64
}

func (e employee) RaiseSalary(f float64) float64 {
	return e.salary * (1 + f)
}

type engine interface {
	start()
	stop()
}

type car struct {
	engine
	wheelCount int
}

func (car *car) start() {
	fmt.Printf("车辆运行\n")
}

func (car *car) stop() {
	fmt.Printf("车辆停止\n")
}

func (car *car) GoToWorkIn() {
	car.start()
	car.stop()
}

func (car *car) numberOfWheels() {
	fmt.Printf("车辆轮子个数为：%d\n", car.wheelCount)
}

type Mercedes struct {
	car
}

func (m *Mercedes) sayHiToMerkel() {
	fmt.Printf("Hi\n")
}

type Point3 struct {
	x, y, z float64
}

func (p *Point3) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y + p.z + p.z)
}

type Base struct {
	id int
}

func (b *Base) Id() int {
	return b.id
}
func (b *Base) SetId(i int) {
	b.id = i
}

type Person2 struct {
	Base
	firstName string
	LastName  string
}

type Employee2 struct {
	Person2
	salary int
}

const LIMIT = 5

type Stack struct {
	stack [LIMIT]int
	at    int
}

func (s *Stack) Push(a int) {
	s.stack[s.at] = a
	s.at += 1
}

func (s *Stack) Pop() int {
	s.at -= 1
	res := s.stack[s.at]
	s.stack[s.at] = 0
	return res
}
