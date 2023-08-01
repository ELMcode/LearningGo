package main

import (
	"fmt"
	"math/rand"
) // pour importer plus simplement plusieurs packages

func main() {
	fmt.Println("My favorite number is", rand.Intn(10)) // rand.Intn provient du module math/rand importé
}
