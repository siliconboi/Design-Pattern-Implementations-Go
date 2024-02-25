package main

import "fmt"

type Car interface {
	Drive(s *Specs)
}
type Specs struct {
	HP       int
	Engine   string
	TopSpeed int
	Distance int
	Company  Car
}

type BMW struct {
}

func (b *BMW) Drive(s *Specs) {
	fmt.Println("Driving a BMW")
}

type Audi struct {
}

func (a *Audi) Drive(s *Specs) {
	fmt.Println("Driving an Audi")
}

type Porsche struct {
}

func (p *Porsche) Drive(s *Specs) {
	fmt.Println("Driving a Porsche")
}

func initSpecs(c Car) *Specs {
	return &Specs{
		HP:       500,
		Engine:   "V8",
		TopSpeed: 200,
		Distance: 0,
	}
}

func (s *Specs) addDistance(d int) {
	s.Distance += d
	fmt.Println("Distance:", s.Distance)
	s.Drive()
}

func (s *Specs) setCar(c Car) {
	s.Company = c
}

func (s *Specs) Drive() {
	s.Company.Drive(s)
}
