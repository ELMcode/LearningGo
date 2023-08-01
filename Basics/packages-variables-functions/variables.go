package main

import "fmt"

var c, python, java bool // déclaration de plusieurs variables avec le même type (bool)

func main() {
	var i int
	fmt.Println(i, c, python, java)
}

// résultat console : 0 false false false
// car i est un int donc 0 par defaut, puis les var c python et java sont des boolean donc false par defaut
