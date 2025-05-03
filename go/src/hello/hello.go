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
}
