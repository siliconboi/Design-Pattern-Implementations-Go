package main

import "fmt"

type Driver struct {
}

func (d *Driver) DriveVehicle(v Vehicle) {
	v.Drive()
}

type Vehicle interface {
	Drive()
}

type Car struct {
}

func (c *Car) Drive() {
	fmt.Println("Driving a car")
}

type Plane struct {
}

func (p *Plane) Fly() {
	fmt.Println("Flying a plane")
}

type PlaneAdapter struct {
	plane *Plane
}

func (pa *PlaneAdapter) Drive() {
	fmt.Println("Adapter converts plane instructions to car instructions")
	pa.plane.Fly()
}
