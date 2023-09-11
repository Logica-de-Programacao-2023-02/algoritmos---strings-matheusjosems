package main

import "fmt"

func main() {
	var s1, s2 string

	fmt.Println("Digite uma string ", s1)
	fmt.Scan(&s1)
	fmt.Println("digite outra string", s2)
	fmt.Scan(&s2)

	fmt.Println(s1 + ", " + s2)
}
