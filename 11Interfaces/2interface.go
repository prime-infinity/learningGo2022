package main

import "fmt"

type ninjaSword struct {
	owner string
}
type ninjaStar struct {
	owner string
}

func (ns ninjaSword) attack() {
	fmt.Println("attacking with sword")
}
func (ns ninjaStar) attack() {
	fmt.Println("attacking with star")
}

type weapon interface {
	attack()
}

func attackAll(w weapon) {
	fmt.Println("attacking all")
	w.attack()
}

func main() {
	weapons := []weapon{ninjaStar{owner: "roshi"}, ninjaSword{owner: "ugwee"}}
	for _, v := range weapons {
		v.attack()
	}
}
