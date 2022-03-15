package main

import "fmt"

func main() {

	ford := &car{"Fiesta", "Ford", 2008}
	merc := &car{"D600", "Mercedes", 2009}
	bmw2 := &car{"X 800", "BMW", 2008}
	// query:
	allCars := cars{ford, bmw2, merc}
	allNewBMWs := allCars.FindAll(func(car *car) bool {
		return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
	})
	fmt.Println("AllCars: ", allCars)
	fmt.Println("New BMWs: ", allNewBMWs)
	//
	manufacturers := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"}
	sortedAppender, sortedCars := MakeSortedAppender(manufacturers)

	//将之前定义的车辆添加至对应品牌
	allCars.Process(sortedAppender)
	fmt.Println("Map sortedCars: ", sortedCars)
	BMWCount := len(sortedCars["BMW"])
	fmt.Println("We have ", BMWCount, " BMWs")
}

type car struct {
	Model        string
	Manufacturer string
	BuildYear    int
}

type cars []*car

type Any interface{}

func (cs cars) Process(f func(c *car)) {
	for _, v := range cs {
		f(v)
	}
}

func (cs cars) FindAll(f func(c *car) bool) cars {
	res := make(cars, 0)
	cs.Process(func(c *car) {
		if f(c) {
			res = append(res, c)
		}
	})
	return res
}

func (cs cars) Map(f func(c *car) Any) []Any {
	res := make([]Any, 0)
	ix := 0
	cs.Process(func(c *car) {
		res[ix] = f(c)
		ix++
	})
	return res
}

func MakeSortedAppender(manufacturer []string) (func(c *car), map[string]cars) {
	sortedCar := make(map[string]cars)
	for _, v := range manufacturer {
		sortedCar[v] = make(cars, 0)
	}
	sortedCar["Default"] = make(cars, 0)
	appender := func(c *car) {
		if _, ok := sortedCar[c.Manufacturer]; ok {
			sortedCar[c.Manufacturer] = append(sortedCar[c.Manufacturer], c)
		} else {
			sortedCar["Default"] = append(sortedCar["Default"], c)
		}
	}
	return appender, sortedCar
}
