package main

import (
	"bufio"   //lecture efficace de l’entrée utilisateur (console).
	"fmt"     //affichage à l’écran et formatage (Println, Printf). Affichage et lecture simple
	"os"      //accès aux fonctionnalités système, ici pour stdin et fichiers. Accès au système (stdin, fichiers, variables d’environnement)
	"strings" //manipulation de chaînes de caractères (TrimSpace, ToLower).
)

func main() {
	reader := bufio.NewReader(os.Stdin) // Création d'un lecteur pour récupérer l'entrée utilisateur

	fmt.Print("Quel est ton nom ? ")
	nom, _ := reader.ReadString('\n') // Lit la ligne entrée par l'utilisateur
	nom = strings.TrimSpace(nom)      // Supprime le saut de ligne

	fmt.Println("Bonjour,", nom, "! Bienvenue dans le monde de Go 🎉")
}
