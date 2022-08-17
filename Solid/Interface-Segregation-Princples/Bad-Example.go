package main

import "fmt"

type Cars interface {
	Run()
	OpenSunRoof()
	EnterPitStop()
}

type FormulaCars struct {
	name    string
	sunroof bool
}
type SportCars struct {
	name    string
	sunroof bool
}

func (f FormulaCars) Run() {
	fmt.Println("started car")
}
func (f FormulaCars) EnterPitStop() {
	fmt.Println("Changed wheel")
}
func (f FormulaCars) OpenSunRoof() {
}

func (s SportCars) Run() {
	fmt.Println("started car")
}
func (s SportCars) EnterPitStop() {
}
func (s SportCars) OpenSunRoof() {
	fmt.Println("Opened Sun Roof")
}

func main() {
	var formulaCars Cars
	formulaCars = FormulaCars{
		name:    "Vettel Cars",
		sunroof: false,
	}

	formulaCars.Run()
	formulaCars.EnterPitStop()

	var sportCars Cars
	sportCars = SportCars{
		name:    "BMW",
		sunroof: true,
	}

	formulaCars.Run()
	sportCars.OpenSunRoof()

}
