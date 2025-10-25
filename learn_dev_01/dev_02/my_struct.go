package main

import "fmt"

type Emploee struct {
	Name string
	ID   int
}

func (e Emploee) description() {
	fmt.Printf("%s: %d\n", e.Name, e.ID)
}

type Manager struct {
	Emploee
	Reports []Emploee
}

func main() {
	emp1 := Emploee{
		Name: "Vasia",
		ID:   12,
	}

	man1 := Manager{
		Emploee: Emploee{
			Name: "Petia",
			ID:   7,
		},
		Reports: []Emploee{
			{Name: "Olia", ID: 3},
			{Name: "Kolya", ID: 6},
		},
	}

	fmt.Println(emp1.Name)
	fmt.Println(man1.Name)
	fmt.Println(man1.Emploee.Name)
	man1.description()
}
