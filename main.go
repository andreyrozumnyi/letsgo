package main

import "fmt"

type human struct {
	Age int
}

type worker struct {
	human
	YearsOfExperience int
	Salary            int
	Company           string
}

type softwareEngineer struct {
	worker
	Stack []string `json:"stack"`
}

type teacher struct {
	worker
	Subject string `json:"subject"`
}

type driver struct {
	worker
	Category string `json:"license_category"`
}

type lawyer struct {
	worker
	Specialty string `json:"specialty"`
}

type doctor struct {
	worker
	Specialty string `json:"specialty"`
}

func main() {
	peter := &softwareEngineer{
		worker: worker{
			YearsOfExperience: 5,
		},
		Stack: []string{"go", "node.js"},
	}
	fmt.Println(peter.YearsOfExperience)

	john := &lawyer{
		worker: worker{
			YearsOfExperience: 3,
		},
		Specialty: "Corporate",
	}

	// will cause panic if uncommented
	// john = nil

	fmt.Println(john.YearsOfExperience)
}
