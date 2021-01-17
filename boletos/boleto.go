package boletos

import (
	"fmt"
	"strconv"
	"strings"

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

func getFatorVencimento(dataVencimento string, zerarVencimento bool) string {
	fatorVencimento := utils.CalcularFatorVencimento(dataVencimento, zerarVencimento)
	return fmt.Sprintf("%04d", fatorVencimento)
}

func getValorBoleto(valor float64) string {
	return utils.FormatarValorBoleto(valor, 10)
}

func getNossoNumero(nossoNumero string, tamanho int) string {
	return utils.FormatarNossoNumero(nossoNumero, tamanho)
}

func getCalculoBaseCodigoBarras(valor string, min int, max int) int {
	return utils.CalcularBaseCodigoBarras(valor, min, max)
}

func getCalculoBaseLinhaDigitavel(valor string, min int, max int, somarAcima bool) int {
	return utils.CalcularBaseLinhaDigitavel(valor, min, max, somarAcima)
}

func getLinhaDigitavel(codigoBarras string, min int, max int, somarAcima bool) LinhaDigitavel {

	var dv, valorBase int
	var linhaDigitavel, linhaDigitavelFormatada strings.Builder

	// 1° campo
	linhaDigitavel.WriteString(codigoBarras[0:4])
	linhaDigitavel.WriteString(codigoBarras[19:20])
	linhaDigitavel.WriteString(codigoBarras[20:24])
	valorBase = getCalculoBaseLinhaDigitavel(linhaDigitavel.String(), 1, 2, true)
	dv = utils.CalcularMod(valorBase, 10, 10)
	linhaDigitavel.WriteString(strconv.Itoa(dv))

	// 2° campo
	linhaDigitavel.WriteString(codigoBarras[24:34])
	valorBase = getCalculoBaseLinhaDigitavel(linhaDigitavel.String()[10:20], 1, 2, true)
	dv = utils.CalcularMod(valorBase, 10, 10)
	linhaDigitavel.WriteString(strconv.Itoa(dv))

	// 3° campo
	linhaDigitavel.WriteString(codigoBarras[34:44])
	valorBase = getCalculoBaseLinhaDigitavel(linhaDigitavel.String()[21:31], 1, 2, true)
	dv = utils.CalcularMod(valorBase, 10, 10)
	linhaDigitavel.WriteString(strconv.Itoa(dv))

	// 4° campo
	linhaDigitavel.WriteString(codigoBarras[4:5])

	// 5° campo
	linhaDigitavel.WriteString(codigoBarras[5:9])
	linhaDigitavel.WriteString(codigoBarras[9:19])

	linhaDigitavelFormatada.WriteString(linhaDigitavel.String()[0:5] + ".")
	linhaDigitavelFormatada.WriteString(linhaDigitavel.String()[5:10] + " ")
	linhaDigitavelFormatada.WriteString(linhaDigitavel.String()[10:15] + ".")
	linhaDigitavelFormatada.WriteString(linhaDigitavel.String()[15:21] + " ")
	linhaDigitavelFormatada.WriteString(linhaDigitavel.String()[21:26] + ".")
	linhaDigitavelFormatada.WriteString(linhaDigitavel.String()[26:32] + " ")
	linhaDigitavelFormatada.WriteString(linhaDigitavel.String()[32:33] + " ")
	linhaDigitavelFormatada.WriteString(linhaDigitavel.String()[33:47])

	linha := LinhaDigitavel{
		LinhaDigitavel:          linhaDigitavel.String(),
		LinhaDigitavelFormatada: linhaDigitavelFormatada.String(),
	}

	return linha
}
