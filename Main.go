package main

import (
	"fmt"

	conta "github.com/YuriBertoldi/Go-ContaCorrente/Models"
)

func main() {
	fmt.Println("teste")
	ccTitularTeste := new(conta.ContaCorrente)
	ccTitularTeste.NumeroAgencia = 123
	ccTitularTeste.NumeroConta = 5123123
	ccTitularTeste.Titular = "Titular teste 1"
	ccTitularTeste.Saldo = 102.45

	ccTitularTeste1 := conta.ContaCorrente{Titular: "Titular teste 3", NumeroAgencia: 321, NumeroConta: 023456, Saldo: 200}

	fmt.Println(ccTitularTeste, " saque: ", ccTitularTeste.Sacar(100), " saldo disp: ", ccTitularTeste.Saldo)

	fmt.Println(ccTitularTeste1, " Transferencia: ", ccTitularTeste1.Transferir(200, ccTitularTeste), " saldo disp: ", ccTitularTeste.Saldo)
}
