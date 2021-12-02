package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Println("Within the save method")
	fmt.Printf("Saving user data %s, with age %d to the database\n", u.name, u.age)
}

func (u user) ageOfmajority() bool {
	return u.age >= 18
}

func (u *user) makeBirthday() {
	u.age++
}

func main() {
	fmt.Println("------------ Methods -----------")
	user1 := user{"User 1", 20}
	user2 := user{"User 2", 25}
	fmt.Println(user1)
	fmt.Println(user2)
	user1.save()
	user2.save()
	ageOfmajority := user2.ageOfmajority()
	fmt.Println(ageOfmajority)
	user2.makeBirthday()
	fmt.Println(user2.age)
}
