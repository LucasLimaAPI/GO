package main // Sempre será necessário importar o pacote

import (
	"fmt"
)

func main() {
	name := "Lucas" //ao invés de usarmos var podemos usar ":="
	version := 1.1
	old := 23

	fmt.Println("Hello Mr.", name, "Are You:", old, "Years old") //Importa a funcionalidade e usa ela novamente dando o nome da function
	fmt.Println("This Program Version is:", version)

	fmt.Println("1- Initialize Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")

	var comand int
	//fmt.Scanf("%d", &comand) //captura o que vai ser escrito pelo usuario "%d" é o modificador, necessário para indicar o que quer receber neste caso um inteiro o & serve para atrelar a variavel a funcao como a usada foi int então todos os numeros inseridos seriam igual a 0, por isso utilizamos &
	fmt.Scan(&comand)
	fmt.Println("Comand:", comand)
}
