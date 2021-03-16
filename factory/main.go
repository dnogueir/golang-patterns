package main

import (
	"fmt"
)

func main() {
	gun1 := createGun("ak47")
	gun2 := createGun("musket")

	fmt.Println(gun1)
	fmt.Println(gun2)
}
