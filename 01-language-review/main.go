package main

import "fmt"

var y int // package level scope

type person struct {
	fname string
	lname string
}

func (p person) speak() {
	fmt.Println(p.fname, `say "Good morning, James"`)
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, "I have licenseToKill:", sa.licenseToKill)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	x := 7 // block level scope
	fmt.Println(x)
	fmt.Printf("%T \n", x)

	y = 44
	fmt.Println(y)

	// composite literal
	xi := []int{1, 2, 3}
	fmt.Println(xi)

	// map
	m := map[string]int{
		"user1": 23,
		"user2": 76,
	}
	fmt.Println(m)

	// struct
	p1 := person{
		"Tony",
		"Stark",
	}
	fmt.Println(p1)
	p1.speak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	sa1.speak()
	sa1.person.speak()

	// polymorphism
	saySomething(p1)
	saySomething(sa1)
}
