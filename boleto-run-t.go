package main

import (
	"fmt"
	"log"

	"github.com/polaris-boleto/boletos"
)

func main() {

	var codBanco string
	var dataVenc string
	var zerarVenc bool
	var valor float64
	var zerarValor bool
	var carteira int16
	var banco boletos.Banco
	var codBeneficiario int32
	var ag int32
	var nossoNumero string

	var cod boletos.CodigoBarras
	var linha boletos.LinhaDigitavel
	var erro boletos.Erro

	/*
		Bradesco - testes
	*/

	codBanco = "237"
	valor = 1248.79
	dataVenc = "1997-10-09"
	carteira = 26
	nossoNumero = "30503154617"
	codBeneficiario = 5331
	zerarVenc = false
	zerarValor = false
	ag = 3381

	// fmt.Printf("Cod: %s - Valor R$: %.2f - Agencia: %d - Cart: %d - Cod Ben: %d - Número: %s\n\n", codBanco, valor, ag, carteira, codBeneficiario, nossoNumero)

	banco, erro = boletos.InstanciarBoleto(codBanco)

	if erro.Status != 0 {
		log.Fatal("Simulando Erro 1 API: ", erro.Mensagem)
	}

	banco.SetCodigo()
	banco.SetAgencia(ag)
	banco.SetCarteira(carteira)
	banco.SetFatorVencimento(dataVenc, zerarVenc)
	banco.SetValorBoleto(valor, zerarValor)
	banco.SetNossoNumero(nossoNumero)
	banco.SetCodigoBeneficiario(codBeneficiario)

	/*
		Santander - testes
	*/

	// codBanco = "033"
	// valor = 80.55
	// dataVenc = "2012-11-26"
	// carteira = 104
	// nossoNumero = "0000000000027"
	// codBeneficiario = int32(5276543)

	// banco = boletos.InstanciarBoleto(codBanco)
	// banco.SetCodigo()
	// banco.SetFatorVencimento(dataVenc)
	// banco.SetValorBoleto(valor)
	// banco.SetCodigoBeneficiario(codBeneficiario)
	// banco.SetNossoNumero(nossoNumero)
	// banco.SetCarteira(carteira)

	cod, erro = banco.GetCodigoBarras()

	if erro.Status != 0 {
		log.Fatal("Erro 1: ", erro.Mensagem)
	}

	fmt.Println("Codigo Barras -> ", cod.CodigoBarrasCompleto)

	linha, erro = banco.GetLinhaDigitavel(cod.CodigoBarrasCompleto)

	if erro.Status != 0 {
		log.Fatal("Erro 2: ", erro.Mensagem)
	}

	fmt.Println("Linha Digitavel", linha.LinhaDigitavel)
	fmt.Println("Linha Digitavel Formatada", linha.LinhaDigitavelFormatada)

}
