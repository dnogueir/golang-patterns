package main

import (
	"fmt"
)

// create an interface with the volatile methods in your code
type doSpecific interface {
	execute()
}

// create struct types that implements the interface

type fastDriving struct {
}

func (f fastDriving) execute() {
	fmt.Println("driving fast")
}

type slowDriving struct {
}

func (s slowDriving) execute() {
	fmt.Println("driving slow")
}

// --------------
// create struct with the main type and a attribute that points to the interface
type vehicle struct {
	name       string
	doSpecific doSpecific
}

func main() {
	// now every value of the type vehicle can have a different behaviour with the doSomething interface

	v1 := vehicle{
		name:       "Celtinha",
		doSpecific: slowDriving{},
	}

	v2 := vehicle{
		name:       "Corolla",
		doSpecific: fastDriving{},
	}

	v1.doSpecific.execute()
	v2.doSpecific.execute()
}
