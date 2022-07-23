package contas

type Conta interface {
   Sacar(value float64) string
}
