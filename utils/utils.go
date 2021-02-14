package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

const formatDate = "2006-01-02"

// CalcularFatorVencimento - calcula o fator de vencimento do boleto
func CalcularFatorVencimento(data string, zerarVencimento bool) int16 {

	var err error

	if zerarVencimento {
		return int16(0)
	}

	dataVencimento, err := time.Parse(formatDate, data)

	if err != nil {
		log.Println(fmt.Sprintf("Valor %v informado para a data de vencimento é inválida. Deve ser informado no formato yyyy-MM-dd.", data))
	}

	dataBase, _ := time.Parse(formatDate, "1997-10-07")

	if dataVencimento.Before(dataBase) {
		log.Println(fmt.Sprintf("Data %v é inválida pois é inferior a data base 1997-10-07.", data))
		return int16(0)
	}

	dias := dataVencimento.Sub(dataBase).Hours() / 24

	return int16(dias)
}

// FormatarValorBoleto - formata o valor do boleto preenchendo com zeros a esquerda
func FormatarValorBoleto(valor float64, tamanho int, zerarValor bool) string {

	var valorBoleto string

	if zerarValor {
		return strings.Repeat("0", tamanho)
	}

	valorBoleto = fmt.Sprintf("%.2f", valor)
	valorBoleto = strings.Replace(valorBoleto, ".", "", -1)
	valorBoleto = strings.Repeat("0", tamanho-len(valorBoleto)) + valorBoleto
	return valorBoleto
}

// FormatarNossoNumero - formata o nosso número preenchendo com zeros a esquerda
func FormatarNossoNumero(nossoNumero string, tamanho int) string {

	if len(nossoNumero) > tamanho {
		return nossoNumero
	}

	return strings.Repeat("0", tamanho-len(nossoNumero)) + nossoNumero
}

// CalcularBaseCodigoBarras - realiza o calculo de base para código de barras
func CalcularBaseCodigoBarras(codigo string, min int, max int) int {

	indice := len(codigo)
	multiplicador := min
	var digito int
	var total int

	for indice >= 1 {
		digito, _ = strconv.Atoi(string(codigo[indice-1]))
		total += digito * multiplicador
		indice--
		multiplicador++
		if multiplicador > max {
			multiplicador = min
		}
	}

	return total
}

// CalcularBaseLinhaDigitavel - realiza o calculo de base para linha digitavel
func CalcularBaseLinhaDigitavel(codigo string, min int, max int, somarAcima bool) int {

	indice := len(codigo)
	multiplicador := max
	var digito int
	var total int

	for indice >= 1 {
		digito, _ = strconv.Atoi(string(codigo[indice-1]))
		totalParcial := (digito * multiplicador)

		if totalParcial > 9 && somarAcima {
			valor := strconv.Itoa(totalParcial)
			digito1, _ := strconv.Atoi(string(valor[0]))
			digito2, _ := strconv.Atoi(string(valor[1]))

			total += (digito1 + digito2)
		} else {
			total += totalParcial
		}

		indice--
		multiplicador--
		if multiplicador < min {
			multiplicador = max
		}
	}

	return total
}

// CalcularMod - realiza calculo da função mod
func CalcularMod(valor, valorSubtrair, mod int) int {
	return (valorSubtrair - (valor % mod))
}
