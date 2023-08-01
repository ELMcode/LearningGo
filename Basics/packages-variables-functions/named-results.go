package main

import "fmt"

//	fonction "split" qui prend un entier 'int' "sum" en entrée et retourne deux entiers "x" et "y".
//
// La fonction utilise une syntaxe spéciale pour retourner deux valeurs à la fois : "x" et "y".
func split(sum int) (x, y int) {

	x = sum * 4 / 9

	y = sum - x

	// La fonction retourne automatiquement les valeurs de "x" et "y" à la fin sans avoir à les spécifier explicitement.
	// On appel ca naked return
	return
}

// La fonction "main" est le point de départ du programme exécutable.
// Elle ne prend pas d'arguments en entrée et ne retourne rien.
func main() {
	// Appelle la fonction "split" avec l'argument 17.
	// La déclaration x, y := est une déclaration d'affectation multiple qui récupère ces deux valeurs renvoyées par split et les attribue aux variables x et y
	x, y := split(17)

	fmt.Println(x, y)
}
