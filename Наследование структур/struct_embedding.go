package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Println(a.Name, "makes a sound")
}

type Dog struct {
	Animal // встраивание
	Breed  string
}

func (d Dog) Bark() {
	fmt.Println(d.Name, "barks")
}

func main() {
	d := Dog{
		Animal: Animal{Name: "Rex"},
		Breed:  "Shepherd",
	}
	d.Speak() // метод Animal доступен напрямую
	d.Bark()
}

// ----------------------------------------------------------------
// Переопределение метода
func (d Dog) Speak() {
	fmt.Println(d.Name, "barks loudly")
}
