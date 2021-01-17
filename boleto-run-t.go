package main

import (
	"fmt"

	"github.com/polaris-boleto/boletos"
)

func main() {

	var codBanco string
	var dataVenc string
	var zerarVenc bool
	var valor float64
	// var zerarValor bool
	var carteira int8
	var cod boletos.CodigoBarras
	var banco boletos.Banco
	var codBeneficiario int32
	var ag int32
	var nossoNumero string

	/*
		Bradesco - testes
	*/

	codBanco = "237"
	dataVenc = "1997-10-02"
	zerarVenc = true
	valor = 1248.79
	// zerarValor = false
	carteira = 26
	ag = 3381
	codBeneficiario = 5331
	nossoNumero = "503154617"

	fmt.Printf("Cod: %s - Valor R$: %.2f - Agencia: %d - Cart: %d - Cod Ben: %d - Número: %s\n\n", codBanco, valor, ag, carteira, codBeneficiario, nossoNumero)

	banco = boletos.InstanciarBoleto(codBanco)
	banco.SetCodigo()
	banco.SetAgencia(ag)
	banco.SetCarteira(carteira)
	banco.SetFatorVencimento(dataVenc, zerarVenc)
	banco.SetValorBoleto(valor)
	banco.SetNossoNumero(nossoNumero)
	banco.SetCodigoBeneficiario(codBeneficiario)

	/*
		Santander - testes
	*/

	// codBanco = "033"
	// valor = 80.55
	// carteira = 104
	// dataVenc = "2012-11-26"
	// nossoNumero = "0000000000027"
	// codBeneficiario = int32(5276543)

	// fmt.Printf("\nCod: %s - Valor R$: %.2f - Cart: %d - Cod Ben: %d - Número: %s\n\n", codBanco, valor, carteira, codBeneficiario, nossoNumero)

	// banco = boletos.InstanciarBoleto(codBanco)
	// banco.SetCodigo()
	// banco.SetFatorVencimento(dataVenc)
	// banco.SetValorBoleto(valor)
	// banco.SetCodigoBeneficiario(codBeneficiario)
	// banco.SetNossoNumero(nossoNumero)
	// banco.SetCarteira(carteira)

	cod = banco.GetCodigoBarras()

	// fmt.Printf("Cod: %s - Valor R$: %s - Campo Livre: %s\n\n", cod.CodigoBanco, cod.Valor, cod.CampoLivre)

	fmt.Println("Codigo Barras", cod.CodigoBarrasCompleto)

	linha := banco.GetLinhaDigitavel(cod.CodigoBarrasCompleto)

	fmt.Println("Linha Digitavel", linha.LinhaDigitavel)
	fmt.Println("Linha Digitavel Formatada", linha.LinhaDigitavelFormatada)

}
