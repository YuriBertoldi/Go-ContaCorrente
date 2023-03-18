package modelConta

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

// func (c *ContaCorrente)  -- igual declaraçao do pascal que identifica de qual obj é
func (c *ContaCorrente) Sacar(vrSaque float64) string {
	if vrSaque > 0 && c.Saldo >= vrSaque {
		c.Saldo -= vrSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(vrDeposito float64) string {
	if vrDeposito > 0 {
		c.Saldo += vrDeposito
		return "Deposito realizado com sucesso"
	} else {
		return "Não é possivel realizar um deposito com valor <= 0"
	}
}

func (c *ContaCorrente) Transferir(vrTransferencia float64, cDestino *ContaCorrente) string {
	if (vrTransferencia > 0) && (c.Saldo >= vrTransferencia) {
		c.Saldo -= vrTransferencia
		cDestino.Depositar(vrTransferencia)
		return "Transferencia realizado com sucesso"
	} else {
		return "Não é possivel realizar um transferencia com valor <= 0"
	}
}
