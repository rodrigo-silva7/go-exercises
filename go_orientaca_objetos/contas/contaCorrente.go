package contas 

import "go_oo/clientes"

type ContaCorrente struct {
   Titular         clientes.Titular
   NumeroAgencia   int
   NumeroConta     int
   saldo           float64
} 

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
   if valorDoSaque > c.saldo && valorDoSaque > 0. {
      return "Não foi possível sacar. saldo insuficiente!";
   } else {
      c.saldo -= valorDoSaque
      return "Saque realizado com sucesso!";
   }
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
   if valorDoDeposito <= 0. {
      return "Não foi possível depositar. Valor zerado ou nulo!", c.saldo;
   } else {
      c.saldo += valorDoDeposito;
      return "Depósito realizado com sucesso!", c.saldo;
   }
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
   if valor <= c.saldo && valor > 0. {
      c.saldo -= valor;
      contaDestino.Depositar(valor);
      return true;
   } else {
      return false;
   }
}

func (c *ContaCorrente) Saldo() float64 {
   return c.saldo;
}
