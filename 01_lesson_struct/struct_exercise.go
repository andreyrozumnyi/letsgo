package main

type Human struct {
	Age int
}

type Worker struct {
	Human
	YearsOfExperience int
	Salary            int
	Company           string
}

type SoftwareEngineer struct {
	Worker
	MainTechnology string `json:"main_technology"`
}

type Teacher struct {
	Worker
	Subject string `json:"subject"`
}

type Driver struct {
	Worker
	Category string `json:"license_category"`
}

type Lawyer struct {
	Worker
	Specialty string `json:"specialty"`
}

type Doctor struct {
	Worker
	Specialty string `json:"specialty"`
}

func main() {

}
