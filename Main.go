package main

import (
	"fmt"

	"github.com/YuriBertoldi/Go-Exemplo-ContaCorrente/Models/modelConta"
)

func main() {
	ccTitularTeste := new(modelConta.ContaCorrente)
	ccTitularTeste.numeroAgencia = 123
	ccTitularTeste.numeroConta = 5123123
	ccTitularTeste.titular = "Titular teste 1"
	ccTitularTeste.Saldo = 102.45

	// ccTitularTeste2 := ContaCorrente{titular: "Titular teste 2", numeroAgencia: 321, numeroConta: 123456, Saldo: 99.89}

	ccTitularTeste1 := modelConta.ContaCorrente{"Titular teste 3", 321, 023456, 200}

	fmt.Println(ccTitularTeste, " saque: ", ccTitularTeste.Sacar(100), " saldo disp: ", ccTitularTeste.Saldo)
	// fmt.Println(ccTitularTeste2, " saque: ", ccTitularTeste2.sacar(100), " saldo disp: ", ccTitularTeste2.Saldo)
	// fmt.Println(ccTitularTeste1, " Deposito: ", ccTitularTeste1.depositar(300), " saldo disp: ", ccTitularTeste1.Saldo)

	fmt.Println(ccTitularTeste1, " Transferencia: ", ccTitularTeste1.Transferir(200, ccTitularTeste), " saldo disp: ", ccTitularTeste.Saldo)
}
