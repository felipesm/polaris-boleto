package boletos

import (
	"fmt"

	"github.com/polaris-boleto/utils"
)

// InstanciarBoleto - retorna a instância de um boleto
func InstanciarBoleto(codigo string) Banco {

	switch {
	case codigo == "237":
		return &Bradesco{}
	case codigo == "033":
		return &Santander{}
	default:
		return nil
	}
}

func retornarFatorVencimento(dataVencimento string) string {
	fatorVencimento := utils.CalcularFatorVencimento(dataVencimento)
	return fmt.Sprintf("%04d", fatorVencimento)
}

func getFatorVencimento(dataVencimento string) string {
	fatorVencimento := utils.CalcularFatorVencimento(dataVencimento)
	return fmt.Sprintf("%04d", fatorVencimento)
}

func getValorBoleto(valor float64) string {
	return utils.FormatarValorBoleto(valor, 10)
}

func getNossoNumero(nossoNumero string, tamanho int) string {
	return utils.FormatarNossoNumero(nossoNumero, tamanho)
}
