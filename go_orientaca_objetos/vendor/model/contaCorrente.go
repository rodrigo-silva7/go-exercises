package model

type ContaCorrente struct {
   Titular         string
   NumeroAgencia   int
   NumeroConta     int
   Saldo           float64
} 

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
   if valorDoSaque > c.Saldo && valorDoSaque > 0. {
      return "Não foi possível sacar. Saldo insuficiente!";
   } else {
      c.Saldo -= valorDoSaque
      return "Saque realizado com sucesso!";
   }
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
   if valorDoDeposito <= 0. {
      return "Não foi possível depositar. Valor zerado ou nulo!", c.Saldo;
   } else {
      c.Saldo += valorDoDeposito;
      return "Depósito realizado com sucesso!", c.Saldo;
   }
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
   if valor <= c.Saldo && valor > 0. {
      c.Saldo -= valor;
      contaDestino.depositar(valor);
      return true;
   } else {
      return false;
   }
}
