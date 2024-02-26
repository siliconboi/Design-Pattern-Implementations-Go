package main

//Decouples abstractions and implementations
//So that we dont have to make a new class for every combination of the two
//Used here for nested if statements if you want to look at it that way
func main() {
	v := &Verstappen{}
	f := &Ferrari{}
	f.setDriver(v)
	f.Drive()
}
