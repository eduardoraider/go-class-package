package main

import (
	"fmt"
	"github.com/eduardoraider/goclass-package/users"
)

func main() {
	u, err := users.New("Eduardo", 0)
	fmt.Println(u, err)
}
