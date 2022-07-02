package main

import "fmt"

type ContaCorrente struct {
   titular         string
   numeroAgencia   int
   numeroConta     int
   saldo           float64
} 

func main() {

   contaRodrigo := ContaCorrente{titular: "Rodrigo", numeroAgencia: 589, numeroConta: 123456, saldo: 125.50}; 
   contaAline := ContaCorrente{"Aline", 590, 123491, 356.15};
   
   fmt.Println(contaRodrigo);
   fmt.Println(contaAline);

}
