package main

// Alternative to if-else and switch
func main() {
	car := &Audi{}
	specs := initSpecs(car)
	specs.setCar(car)
	specs.addDistance(4)
}
