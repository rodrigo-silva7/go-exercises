package contas

import "go_oo/clientes"

type ContaPoupanca struct {
   Titular clientes.Titular
   NumeroAgencia, NumeroConta, Operacao int
   saldo float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
   if valorDoSaque > c.saldo && valorDoSaque > 0. {
      return "Não foi possível sacar. saldo insuficiente!";
   } else {
      c.saldo -= valorDoSaque
      return "Saque realizado com sucesso!";
   }
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
   if valorDoDeposito <= 0. {
      return "Não foi possível depositar. Valor zerado ou nulo!", c.saldo;
   } else {
      c.saldo += valorDoDeposito;
      return "Depósito realizado com sucesso!", c.saldo;
   }
}

func (c *ContaPoupanca) Transferir(valor float64, contaDestino *ContaPoupanca) bool {
   if valor <= c.saldo && valor > 0. {
      c.saldo -= valor;
      contaDestino.Depositar(valor);
      return true;
   } else {
      return false;
   }
}

func (c *ContaPoupanca) Saldo() float64 {
   return c.saldo;
}
