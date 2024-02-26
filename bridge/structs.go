package main

import "fmt"

type Car interface {
	Drive()
	setDriver(Driver)
}

type Ferrari struct {
	driver Driver
}

func (f *Ferrari) Drive() {
	fmt.Println("Drive Ferrari")
	f.driver.DriveCar()
}

func (f *Ferrari) setDriver(driver Driver) {
	f.driver = driver
}

type RedBull struct {
	driver Driver
}

func (r *RedBull) Drive() {
	fmt.Println("Drive RedBull")
	r.driver.DriveCar()
}

func (r *RedBull) setDriver(driver Driver) {
	r.driver = driver
}

type Driver interface {
	DriveCar()
}

type Verstappen struct{}

func (v *Verstappen) DriveCar() {
	fmt.Println("Driving with Max Verstappen")
}

type Vettel struct{}

func (v *Vettel) DriveCar() {
	fmt.Println("Driving with Sebastian Vettel")
}
