package main

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
}
