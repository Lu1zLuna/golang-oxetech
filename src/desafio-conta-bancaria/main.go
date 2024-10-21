package main

import (
	"fmt"
	"time"
)

func main () {
	// OPerações de saque e depósito, testando os limites impostos para cada conta
/* 	contaCorrente := ContaBancaria{
		Titular: "Luiz",
		NumeroConta: "123456",
		Saldo: 1000,
		TipoConta: "corrente",
		LimiteEspecial: 500, // Permite negativar o saldo até -500 sem causar erro
		LimiteSaquesDia: 0,  // Limite diário para contas do tipo poupança
		SaquesHoje: 0,
		Historico: []Transacao{}
	} 
	
	contaCorrente.Depositar(300)
	contaCorrente.ExibirSaldo()
	contaCorrente.Sacar(1500) // Saque dentro do limite especial: saldo 1000 + limite -500
	contaCorrente.ExibirSaldo()
	contaCorrente.Sacar(200)  // Tentativa de saque além do limite especial
	fmt.Println("-------------------")

	contaPoupanca := ContaBancaria{
		Titular: "Pedro",
		NumeroConta: "987654",
		Saldo: 2000,
		TipoConta: "poupança",
		LimiteEspecial: 0, // Conta poupança não possui um limite especial
		LimiteSaquesDia: 3,
		SaquesHoje: 0,
		Historico: []Transacao{}
	}  
	contaPoupanca.Depositar(200)
	contaPoupanca.Sacar(100)  // Saque permitido
	contaPoupanca.ExibirSaldo()
	contaPoupanca.Sacar(700)  // Tentativa de saque além do saldo
	contaPoupanca.Depositar(1100)
	contaPoupanca.Sacar(100)
	contaPoupanca.ExibirSaldo()
	contaPoupanca.Sacar(500)  // Tentativa ultrapassa o limite de saques diários

	// Reinicia limite de saques diários
	contaPoupanca.ResetarSaquesDiarios()
	contaPoupanca.Sacar(500)
	contaPoupanca.ExibirSaldo()

	fmt.Println("-------------------") */

	// Simulando uma transferência entre contas:
	contaCorrente := ContaBancaria{"Luiz", "123456", 1000, "corrente", 500, 0, 0, []Transacao{}}
	contaPoupanca := ContaBancaria{"João", "654321", 500, "poupança", 0, 0, 3, []Transacao{}}

	// Exibe saldo inicial
	contaCorrente.ExibirSaldo()
	contaPoupanca.ExibirSaldo()

	// Realiza transferência da conta corrente para a conta poupança
	fmt.Println("\nTransferindo 200 da conta corrente para a conta poupança...")
	contaCorrente.TransferirEntreContas(&contaPoupanca, 200)

	// Exibe saldo após a transferência
	contaCorrente.ExibirSaldo()
	contaPoupanca.ExibirSaldo()

	// Tentativa de transferir entre a mesma conta (mesmo número de conta)
	fmt.Println("\nTentativa de transferir entre a mesma conta...")
	contaCorrente.TransferirEntreContas(&contaCorrente, 100)

	// Exibe saldo final
	contaCorrente.ExibirSaldo()
	contaPoupanca.ExibirSaldo()
}

type ContaBancaria struct {
	Titular string
	NumeroConta string
	Saldo float32
	TipoConta string
	LimiteEspecial float32
	LimiteSaquesDia int
	SaquesHoje int
	Historico []Transacao
}

type Transacao struct {
	Tipo string
	Valor float32
	Data string
	ContaDestino string
}

func (c ContaBancaria) ExibirSaldo()  {
	fmt.Printf("Seu saldo atual é: %.2f\n", c.Saldo)
}

func (c *ContaBancaria) Depositar(deposito float32) {
	if deposito <= 0 {
		fmt.Println("Valor selecionado para depósito é inválido. " +
			"Tente novamente.")
		return
	}
		c.Saldo += deposito
		fmt.Printf("Depósito no valor de R$%.2f realizado com sucesso!\n", deposito)
		
		c.RegistrarTransacao("Depósito", deposito, "")
}

func (c *ContaBancaria) Sacar(saque float32) {
	if saque <= 0 {
		fmt.Println("Valor selecionado para saque é inválido. Tente novamente.")
		return
	}

	// Verificar o limite de saques diários para poupança
	if c.TipoConta == "poupança" && c.SaquesHoje >= c.LimiteSaquesDia {
		fmt.Printf("Você atingiu o limite de %d saques diários para sua conta poupança.\n", c.LimiteSaquesDia)
		return
	}

	switch c.TipoConta {
		case "corrente":
			// Para a conta corrente, o saldo pode ficar negativo até o limite especial
			if c.Saldo + c.LimiteEspecial >= saque {
				c.Saldo -= saque
				fmt.Println("Saque realizado com sucesso na conta corrente!")
			} else {
				fmt.Printf("O valor do saque (%.2f) excede o saldo disponível e o limite especial. Saldo disponível: %.2f\n", saque, c.Saldo)
			}

		case "poupança":
			// Para a conta poupança, o saldo nunca pode ser negativo
			if c.Saldo >= saque {
				c.Saldo -= saque
				c.SaquesHoje++
				fmt.Println("Saque realizado com sucesso na conta poupança!")

				c.RegistrarTransacao("Depósito", saque, "")
			} else {
				fmt.Printf("O saldo atual (%.2f) é insuficiente para realizar o saque de %.2f na conta poupança.\n", c.Saldo, saque)
			}

		default:
			fmt.Println("Tipo de conta inválido.")
	}
}

func (c *ContaBancaria) ResetarSaquesDiarios() {
	c.SaquesHoje = 0
	fmt.Println("O contador de saques diários foi resetado.")
}

func (c *ContaBancaria) TransferirEntreContas(destino *ContaBancaria, valor float32) {
	if valor <= 0 {
		fmt.Println("O valor da transferência deve ser positivo.")
		return
	}

	// Verifica se a conta de origem e a conta de destino são diferentes
	if c.NumeroConta == destino.NumeroConta {
		fmt.Println("Transferência inválida: as contas de origem e destino devem ser diferentes.")
		return
	}


	switch c.TipoConta {
	case "corrente":
		// Verifica se há saldo suficiente considerando o limite especial
		if c.Saldo + c.LimiteEspecial >= valor {
			c.Saldo -= valor
			destino.Saldo += valor
			fmt.Printf("Transferência de %.2f realizada com sucesso de %s para %s.\n", valor, c.Titular, destino.Titular)

			// Registra uma transação no histórico de ambas as contas
			c.RegistrarTransacao("Transferência enviada", valor, destino.NumeroConta)
			c.RegistrarTransacao("Transferência recebida", valor, c.NumeroConta)
		} else {
			fmt.Printf("Saldo insuficiente para transferir %.2f. Saldo disponível: %.2f\n", valor, c.Saldo+c.LimiteEspecial)
		}

	case "poupança":
		// Não permite saldo negativo na poupança
		if c.Saldo >= valor {
			c.Saldo -= valor
			destino.Saldo += valor
			fmt.Printf("Transferência de %.2f realizada com sucesso de %s para %s.\n", valor, c.Titular, destino.Titular)
		} else {
			fmt.Printf("Saldo insuficiente para transferir %.2f. Saldo disponível: %.2f\n", valor, c.Saldo)
		}

	default:
		fmt.Println("Tipo de conta inválido.")
	}

}

func (c *ContaBancaria) RegistrarTransacao(tipo string, valor float32, contaDestino string) {
	transacao := Transacao{
		Tipo:         tipo,
		Valor:        valor,
		Data:         time.Now().Format("dd/mm/yyyy hh:mm:ss"),
		ContaDestino: contaDestino,
	}
	c.Historico = append(c.Historico, transacao)
}