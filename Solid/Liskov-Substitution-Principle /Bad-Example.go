package main

import "fmt"

type Student interface {
	getName() string
	getClassNumber() int
	getPhysicsPoint() int
	getArtPoint() int
}

type HighSchoolStudent struct {
	name         string
	classNumber  int
	PhysicsPoint int
	ArtPoint     int
}

type StudentScinece struct {
	HighSchoolStudent
}
type StudentArt struct {
	HighSchoolStudent
}

func (h HighSchoolStudent) getName() string {
	return h.name
}
func (h HighSchoolStudent) getClassNumber() int {
	return h.classNumber
}
func (h HighSchoolStudent) getArtPoint() int {
	return h.ArtPoint
}
func (h HighSchoolStudent) getPhysicsPoint() int {
	return h.PhysicsPoint
}

type Printer struct {
}

func (Printer) printStudentInfo(s Student) {
	fmt.Println("Name:", s.getName())
	fmt.Println("Class Number:", s.getClassNumber())
	fmt.Println("Physics Point", s.getPhysicsPoint())
	fmt.Println("Art Point", s.getArtPoint())
}

func main() {
	h := HighSchoolStudent{
		name:        "Samet",
		classNumber: 12,
	}
	ss := StudentScinece{
		HighSchoolStudent{
			name:         "Kadir",
			classNumber:  9,
			PhysicsPoint: 100,
		},
	}
	sa := StudentArt{
		HighSchoolStudent{
			name:        "Baran",
			classNumber: 10,
			ArtPoint:    100,
		},
	}

	p := Printer{}
	p.printStudentInfo(h)
	p.printStudentInfo(ss)
	p.printStudentInfo(sa)

}
