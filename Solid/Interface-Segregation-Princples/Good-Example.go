package main

import "fmt"

type Cars interface {
	Run()
}

type SunRoof interface {
	OpenSunRoof()
}

type PitStop interface {
	EnterPitStop()
}

type FormulaCars struct {
	name string
	Cars
	PitStop
}

type SportCars struct {
	name string
	Cars
	SunRoof
}

func (f *FormulaCars) Run() {
	fmt.Println("started car")
}
func (f *FormulaCars) EnterPitStop() {
	fmt.Println("Changed wheel")
}

func (s *SportCars) Run() {
	fmt.Println("started car")
}

func (s *SportCars) OpenSunRoof() {
	fmt.Println("Opened Sun Roof")
}

func main() {
	var formulaCars FormulaCars
	formulaCars = FormulaCars{
		name: "Vettel",
	}
	formulaCars.Run()
	formulaCars.EnterPitStop()

	var sportCars SportCars
	
	sportCars = SportCars{
		name: "BMW",
	}

	sportCars.Run()
	sportCars.OpenSunRoof()

}
