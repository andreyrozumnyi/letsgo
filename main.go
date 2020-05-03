package main

import "fmt"

type human struct {
	Age       int
	FirstName string
	LastName  string
}

func (h *human) getFullName() string {
	return h.FirstName + " " + h.LastName
}

func (h *human) introduce() {
	fmt.Println("Hi! I am ", h.getFullName())
}

func (h *human) meditate() {
	fmt.Println("I am", h.getFullName(), "starting to breath slowly...")
}

type worker struct {
	*human
	YearsOfExperience int
	Salary            int
	Company           string
}

type softwareEngineer struct {
	*worker
	Stack []string `json:"stack"`
}

type teacher struct {
	*worker
	Subject string `json:"subject"`
}

type driver struct {
	*worker
	Category string `json:"license_category"`
}

type lawyer struct {
	*worker
	Specialty string `json:"specialty"`
}

type doctor struct {
	*worker
	Specialty string `json:"specialty"`
}

func main() {
	peter := &softwareEngineer{
		worker: &worker{
			YearsOfExperience: 5,
		},
		Stack: []string{"go", "node.js"},
	}
	// will cause panic if uncommented
	// peter = nil
	// peter.getFullName()
	fmt.Println(peter.YearsOfExperience)

	john := &lawyer{
		worker: &worker{
			YearsOfExperience: 3,
			human: &human{
				FirstName: "John",
				LastName:  "Doe",
			},
		},
		Specialty: "Corporate",
	}

	fmt.Println(john.getFullName())
	john.introduce()
	john.meditate()
}
