package main

import (
	"fmt"
)

type contato struct {
	nome     string
	telefone string
	email    string
}

// declaração de variáveis:
var contatos = []contato{}
var flag = true
var input rune // rune == char
var novo contato

func main() {
	for flag {
		var novoNome string
		var novoTelefone string
		var novoEmail string

		fmt.Print(`
			*--------------------*
			|        menu        |
			|                    |
			|  [1] adicionar     |
			|  [2] listar        |
			|                    |
			|  [9] sair          |
			|                    |
			*--------------------*

			`)

		fmt.Println("Insira um numero:")
		fmt.Scanf("%c ", &input)
		switch input {
		case '1':
			fmt.Println("Insira um Nome: ")
			fmt.Scanf("%v", &novoNome)
			novo.nome = novoNome

			fmt.Println("Insira um Numero: ")
			fmt.Scan(&novoTelefone)
			novo.telefone = novoTelefone

			fmt.Println("Insira um Email: ")
			fmt.Scan(&novoEmail)
			novo.email = novoEmail

			contatos = append(contatos, novo)
			fmt.Println("\nNovo contato adicionado...")
		case '2':
			if len(contatos) > 0 {
				for _, v := range contatos {
					fmt.Printf("%s | %s | %s\n", v.nome, v.email, v.telefone)
					fmt.Println()
				}

			} else {
				fmt.Println("A lista está vazia... \nInsira novos contatos para mostrar a lista!")
			}
		case '9':
			// sair do loop e terminar o programa
			flag = false
		default:
			fmt.Println("Digite apenas um numero válido...")
		}
	}
}
