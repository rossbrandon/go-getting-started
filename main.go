package main

import (
	"fmt"
	"github.com/rosstafarian/go-getting-started/models"
)

func main() {
	u := models.User{
		ID: 2,
		FirstName: "Ross",
		LastName: "Brandon",
	}
	fmt.Println(u)
}
