package main

import (
   "fmt"
   "go_oo/contas"
   "go_oo/clientes"
)

func pagarBoleto(conta contas.Conta, value float64) {
   conta.Sacar(value);
}

func main() {

   conta1 := contas.ContaCorrente{
      Titular: clientes.Titular{
         Nome: "Bruno",
         CPF: "123.465.879-87",
         Profissao: "Desenvolvedor",
      },
      NumeroAgencia: 123,
      NumeroConta: 123456,
   }

   conta1.Depositar(500);
   fmt.Println(conta1);
   fmt.Println("O Saldo é", conta1.Saldo());

   pagarBoleto(&conta1, 60);

   fmt.Println("O Saldo é", conta1.Saldo());
}
