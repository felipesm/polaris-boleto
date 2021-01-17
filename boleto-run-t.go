package main

import (
	"fmt"

	"github.com/polaris-boleto/boletos"
)

func main() {

	codBanco := "237"
	// ag := int32(3381)
	valor := 1248.79
	carteira := int8(26)
	codBeneficiario := int32(5331)
	nossoNumero := "503154617"
	var dataVenc string
	var cod boletos.CodigoBarras
	var banco boletos.Banco

	// dataVenc = "2021-01-11"

	// fmt.Printf("Cod: %s - Valor R$: %.2f - Agencia: %d - Cart: %d - Cod Ben: %d - Número: %s\n\n", codBanco, valor, ag, carteira, codBeneficiario, nossoNumero)

	// banco = boletos.InstanciarBoleto(codBanco)
	// banco.SetCodigo()
	// banco.SetAgencia(ag)
	// banco.SetCarteira(carteira)
	// banco.SetFatorVencimento(dataVenc)
	// banco.SetValorBoleto(valor)
	// banco.SetNossoNumero(nossoNumero)
	// banco.SetCodigoBeneficiario(codBeneficiario)

	// cod = banco.GetCodigoBarras()

	// fmt.Printf("Cod: %s - Valor R$: %s - Campo Livre: %s\n\n", cod.CodigoBanco, cod.Valor, cod.CampoLivre)

	// fmt.Println("Codigo Barras Bradesco", cod)

	// cdBarras := "2379843600000432743381260044107515500053310"

	// dv := getDVCodigoBarras(cdBarras)

	// fmt.Println(dv)

	//Santander - testes
	codBanco = "033"
	// ag = int32(2)
	valor = 900.00
	carteira = 101
	dataVenc = "2017-03-04"
	nossoNumero = "0000000000195"
	codBeneficiario = int32(7964234)

	fmt.Printf("\nCod: %s - Valor R$: %.2f - Cart: %d - Cod Ben: %d - Número: %s\n\n", codBanco, valor, carteira, codBeneficiario, nossoNumero)

	banco = boletos.InstanciarBoleto(codBanco)
	banco.SetCodigo()
	banco.SetFatorVencimento(dataVenc)
	banco.SetValorBoleto(valor)
	banco.SetCodigoBeneficiario(codBeneficiario)
	banco.SetNossoNumero(nossoNumero)
	banco.SetCarteira(carteira)

	cod = banco.GetCodigoBarras()

	fmt.Printf("Cod: %s - Valor R$: %s - Campo Livre: %s\n\n", cod.CodigoBanco, cod.Valor, cod.CampoLivre)

	fmt.Println("Codigo Barras Santander", cod.CodigoBarrasCompleto)

	linha := banco.GetLinhaDigitavel(cod.CodigoBarrasCompleto)

	fmt.Println("Linha Digitavel", linha.LinhaDigitavel)
	fmt.Println("Linha Digitavel Formatada", linha.LinhaDigitavelFormatada)

}
