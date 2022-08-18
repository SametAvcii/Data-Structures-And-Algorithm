package main

import "fmt"

type Cars interface {
	Run()
}

type SunRoof interface {
	Cars
	OpenSunRoof()
}

type PitStop interface {
	Cars
	EnterPitStop()
}

type FormulaCars struct {
	name string
}

type SportCars struct {
	name string
}

func (f FormulaCars) Run() {
	fmt.Println("started car")
}
func (f FormulaCars) EnterPitStop() {
	fmt.Println("Changed wheel")
}

func (s SportCars) Run() {
	fmt.Println("started car")
}

func (s SportCars) OpenSunRoof() {
	fmt.Println("Opened Sun Roof")
}

func main() {
	var formulaCars PitStop
	formulaCars = FormulaCars{
		name: "Vettel",
	}

	formulaCars.Run()
	formulaCars.EnterPitStop()

	var sportCars SunRoof

	sportCars = SportCars{
		name: "BMW",
	}

	sportCars.Run()
	sportCars.OpenSunRoof()

}
