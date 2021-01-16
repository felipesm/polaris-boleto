package utils

import (
	"fmt"
	"strings"
	"time"
)

// CalcularFatorVencimento - calcula o fator de vencimento do boleto
func CalcularFatorVencimento(data string) int16 {

	dataVencimento, _ := time.Parse("2006-01-02", data)

	dataBase, _ := time.Parse("2006-01-02", "1997-10-07")

	dias := dataVencimento.Sub(dataBase).Hours() / 24

	return int16(dias)
}

// FormatarValorBoleto - formata o valor do boleto preenchendo com zeros a esquerda
func FormatarValorBoleto(valor float64, tamanho int) string {
	valorBoleto := fmt.Sprintf("%.2f", valor)
	valorBoleto = strings.Replace(valorBoleto, ".", "", -1)
	valorBoleto = strings.Repeat("0", tamanho-len(valorBoleto)) + valorBoleto
	return valorBoleto
}

// FormatarNossoNumero - formata o nosso número preenchendo com zeros a esquerda
func FormatarNossoNumero(nossoNumero string, tamanho int) string {
	return strings.Repeat("0", tamanho-len(nossoNumero)) + nossoNumero
}
