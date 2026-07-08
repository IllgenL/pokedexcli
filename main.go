package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		text = strings.TrimSpace(text)
		words := cleanInput(text)
		fmt.Println("Your command was:", words[0])
	}
}
