package main

import "fmt"
type Manager interface {
	cheek()
	Setnumber()
	Print()
}

type commodity struct {
	name   string
	price  int
	nubmer int
}

func (c commodity) cheek() {
	fmt.Println(c.nubmer)
}
func (c *commodity) Update(a int) {
	c.nubmer = a

}
func (c commodity) Print() {
	fmt.Println(c)

}
type EManger interface {
	Echeak()
	Eupdate()
	EPrint()
}



type electronics struct {
	brand     string
	Model     string
	commodity commodity
}

func (e electronics) Echeak() {
	fmt.Println(e.commodity.nubmer)
}
func (e *electronics) EUpdate(a int) {
	e.commodity.nubmer = a
}
func (e electronics) EPrint() {
	fmt.Println(e)
}


