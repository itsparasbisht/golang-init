package main

import "fmt"

type User struct {
	id         int
	email      string
	item_count int
}

func (u *User) addItem(count int) {
	u.item_count += count
}

func (u *User) removeItem(count int) {
	u.item_count -= count
}

func main() {
	user1 := User{1, "john@mail.com", 0}
	user1.addItem(2)
	user1.addItem(5)
	user1.removeItem(2)
	fmt.Println(user1)
}