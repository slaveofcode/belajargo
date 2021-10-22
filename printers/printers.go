package printers

type Child struct {
	Name string
	age  int
}

func NewChild(name string, age int) Child {
	return Child{
		Name: name,
		age:  age,
	}
}
