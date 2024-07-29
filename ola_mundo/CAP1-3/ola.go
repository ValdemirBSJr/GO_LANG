// para iniciar o gerenciador de dependencias do projeto: go mod init ola_mundo
// para rodar o projeto: go run . ou go run arquivo.go
// pra dizer ao go qual o package principal é o main (onde começa e onde termina. Demais coisas vc importa)
// em go nao pode haver variaveis declaradas e nao usadas
// pra descartar um valor que nao vai ser usado usa _
// https://go.dev/play/
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Olá mundo!")
}
