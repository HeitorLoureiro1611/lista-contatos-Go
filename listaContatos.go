package main

import (
	"bufio"
	"fmt"
	"os"
)

type contato struct {
	nome     string
	telefone string
	email    string
}

func main() {
	// declaração de variáveis:
	// todos os itens estão aqui dentro
	var contatos = []contato{}
	var remove int
	// flag para não ter q lidar com while true puro, assim caso tenha break no codigo ele não quebre
	// apenas sai quando essa variável é alterada
	var flag = true
	var input rune // rune == char

	// Criando um scanner
	// Scanner com tratamento de bufer pra caso eu tenha problemas com espaços nos inputs de string
	sc := bufio.NewScanner(os.Stdin)

	for flag {
		// Novo está declarado aqui para que toda vez que o programa reiniciar ele zerar o conteúdo de dentro dele
		var novo contato

		fmt.Print(`
			*--------------------*
			|        menu        |
			|                    |
			|  [1] adicionar     |
			|  [2] listar        |
			|  [3] Apagar        |
			|                    |
			|  [9] sair          |
			|                    |
			*--------------------*

			`)

		fmt.Print("Insira um numero: ")
		// uso de scanf pela praticidade
		fmt.Scanf("%c ", &input)
		fmt.Println()

		switch input {
		case '1': // rune tem q ser lidado como char então: 'x'
			fmt.Print("Insira um Nome: ")
			sc.Scan()             // o sc.Scan leva o input para o sc.Text ou sc.Bytes para poder ser usado
			novo.nome = sc.Text() // insiro o meu texto digitado na variável

			fmt.Print("Insira um Numero: ")
			sc.Scan()
			novo.telefone = sc.Text()

			fmt.Print("Insira um Email: ")
			sc.Scan()
			novo.email = sc.Text()

			contatos = append(contatos, novo) // adiciona o novo contato para a lista total de contatos
			fmt.Println("\nNovo contato adicionado...")

		case '2':
			if len(contatos) > 0 {
				for k, v := range contatos {
					fmt.Printf("Id: %d | nome: %s | email: %s | numero: %s\n", k+1, v.nome, v.email, v.telefone)
				}

			} else {
				fmt.Println("A lista está vazia... \nInsira novos contatos para mostrar a lista!")
			}

		case '3':
			if len(contatos) == 0 {
				fmt.Println("A lista está vazia... \nInsira novos contatos para mostrar a lista!")

			} else {
				fmt.Print("Insira o ID do contato para ser apagado: ")
				fmt.Scanf("%d\n", &remove)
				// checa se o input está dentro dos limites da lista
				if remove > 0 && remove <= len(contatos) {
					if len(contatos) > 1 {
						// essa logica só funciona quando eu tenho mais de um contato na lista
						remove = remove - 1
						contatos = append(contatos[:remove], contatos[remove+1:]...)
					} else if len(contatos) == 1 && remove == 1 {
						// caso eu tenho 1 só contato e eu quero apagar ele
						contatos = nil
					}
					fmt.Println("Contato deletado...")

				} else {
					fmt.Println("ID inválido...")

				}
			}
		case '9':
			// sair do loop e terminar o programa
			flag = false
		default:
			fmt.Println("Digite apenas um numero válido...")
		}
	}
	fmt.Println("Obrigado por utilizar!")
}
