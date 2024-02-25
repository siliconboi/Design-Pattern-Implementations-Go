package main

// The Adapter acts as a wrapper between two objects.
// It takes two incompatible interfaces and allows them to work together.
// It is similar to Bridge, State and Strategy
func main() {
	Driver := &Driver{}
	car := &Car{}
	Driver.DriveVehicle(car)
	plane := &Plane{}
	planeAdapter := &PlaneAdapter{plane}
	Driver.DriveVehicle(planeAdapter)
}
