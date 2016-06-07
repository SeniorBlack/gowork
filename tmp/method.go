package main

import (
	"fmt"
)

type User struct {
	name string
	age  int
}

func (p *User) kk() {
	get()
}
