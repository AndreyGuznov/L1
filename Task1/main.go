package main

import "fmt"

type Human struct {
	Name     string
	Surname  string
	Sex      string
	JobTitle string
	Salary   int
	Age      int
}

type Action struct {
	Human
}

func (h *Human) IsAdult() bool {
	if h.Age < 18 {
		return false
	}
	return true
}

func (h *Human) Promotion(newJobTitle string) {
	h.JobTitle = newJobTitle
}

func (h *Human) ChangeSalary(newSalary int) {
	h.Salary = newSalary
}

func main() {
	human1 := Human{
		Name:     "Tom",
		Surname:  "Jons",
		Sex:      "Male",
		JobTitle: "ThirdCategory",
		Salary:   1000,
		Age:      31,
	}
	action1 := Action{
		human1,
	}
	res1 := action1.IsAdult()
	action1.Promotion("secondCategori")
	action1.ChangeSalary(2000)
	fmt.Println(res1)
	fmt.Println(action1)
}
