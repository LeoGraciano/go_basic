package main

import "fmt"

type user struct {
	name string
	age  int
}

func (u user) save() {
	fmt.Printf("Salvando %s...\n", u.name)
}
func (u user) majorAge() bool {
	return u.age >= 18
}

func (u *user) makerBirthday() {
	u.age++
}

func write() {
	fmt.Println("Escrevendo do arquivo!")
}

func main() {
	user := user{"Leonardo", 39}
	fmt.Printf("Nome: %s, Idade: %d\n", user.name, user.age)
	user.save()
	major := user.majorAge()
	fmt.Printf("É maior de idade: %t\n", major)
	user.makerBirthday()
	fmt.Printf("Nova idade: %d\n", user.age)

}
