package main

import "fmt"

type Student interface {
	getName() string
	getClassNumber() int
}

type HighSchoolStudent struct {
	name        string
	classNumber int
}

func (h HighSchoolStudent) getName() string {
	return h.name
}
func (h HighSchoolStudent) getClassNumber() int {
	return h.classNumber
}

type StudentScinece struct {
	HighSchoolStudent
	PhysicsPoint int
}
type StudentArt struct {
	HighSchoolStudent
	ArtPoint int
}

type Printer struct {
}

func (Printer) printStudentInfo(s Student) {
	fmt.Println("name:", s.getName())
	fmt.Println("class number:", s.getClassNumber())
}

func main() {
	h := HighSchoolStudent{
		name:        "Samet",
		classNumber: 12,
	}
	ss := StudentScinece{
		HighSchoolStudent: HighSchoolStudent{
			name:        "Kadir",
			classNumber: 9,
		},
		PhysicsPoint: 78,
	}
	sa := StudentArt{
		HighSchoolStudent: HighSchoolStudent{
			name:        "Baran",
			classNumber: 10,
		},
		ArtPoint: 100,
	}

	p := Printer{}
	p.printStudentInfo(h)
	p.printStudentInfo(ss)
	p.printStudentInfo(sa)

}
