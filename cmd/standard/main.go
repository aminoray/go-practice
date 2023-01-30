package main

import (
	"fmt"

	"github.com/FJC-OMUSUBI/go-handson/internal/user/domain"
)

func main() {
	p := domain.User{
		Name:  "モック",
		Email: "abc@yahoo.com",
	}
	fmt.Printf("%+v", p)
}
