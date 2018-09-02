package main

import (
	"fmt"
)

type Animal interface {
	sleep()
	move()
}

type Dog struct {
	name string
}

func (d Dog) sleep() {
	fmt.Println(d.name + " sleepping")
}

func (d Dog) move() {
	fmt.Println(d.name + " moving")
}

func (d Dog) eat() {
	fmt.Println(d.name + " eating")
}

func sleep(animal Animal) {
	animal.sleep()
}

type Cat struct {
	name string
}

func (c Cat) move() {
	fmt.Println(c.name + " moving")
}

func (c Cat) sleep() {
	fmt.Println(c.name + " sleepping")
}

func main() {
	var animal Animal
	animal = Dog{"Dim"}

	switch v := animal.(type) {
	case Dog:
		fmt.Println(v.name, "Dog")
		break
	default:
		break
	}

	dog, ok := animal.(Dog)
	if ok {
		dog.eat()
	}

	cat, ok := animal.(Cat)
	if ok {
		cat.sleep()
	}

	//animal.move()
	//animal.sleep()

	//print:
	//Dim moving
	//Dim sleepping

	//	d := Dog{"Dim"}
	//	sleep(d)

	//print:
	//Dim sleepping
}
