package boletos

import (
	"fmt"
	"strconv"
)

// Bradesco - estrutura para representar o banco Bradesco
type Bradesco struct {
	codigoBradesco    string // 01 a 03 - identificacao do banco
	codigoMoeda       string // 04 a 04 - codigo da moeda (real = 9 | outros = 0)
	dvCodigoBarras    string // 05 a 05 - digito verificador codigo de barras
	fatorVencimento   string // 06 a 09 - fator vencimento
	valor             string // 10 a 19 - valor
	agencia           string // 20 a 23 - agencia (campo livre)
	carteira          string // 24 a 25 - carteira (campo livre)
	nossoNumero       string // 26 a 36 - nosso numero (campo livre)
	contaBeneficiario string // 37 a 43 - conta beneficiario (campo livre)
	zero              string // 44 a 44 - zero
}

const codigoBradesco string = "237"

// SetCodigo - atribuir código boleto Bradesco
func (b *Bradesco) SetCodigo() {
	b.codigoBradesco = codigoBradesco
}

// SetAgencia - atribuir agência boleto Bradesco
func (b *Bradesco) SetAgencia(agencia int32) {
	b.agencia = fmt.Sprintf("%04d", agencia)
}

// SetCarteira - atribuir código carteira boleto Bradesco
func (b *Bradesco) SetCarteira(carteira int8) {
	b.carteira = fmt.Sprintf("%02d", carteira)
}

// SetCodigoBeneficiario - atribuir conta beneficiário boleto Bradesco
func (b *Bradesco) SetCodigoBeneficiario(contaBeneficiario int32) {
	b.contaBeneficiario = fmt.Sprintf("%07d", contaBeneficiario)
}

// SetFatorVencimento - atribuir fator de vencimento boleto Bradesco
func (b *Bradesco) SetFatorVencimento(dataVencimento string) {
	b.fatorVencimento = getFatorVencimento(dataVencimento)
}

// SetValorBoleto - atribuir valor do boleto Bradesco
func (b *Bradesco) SetValorBoleto(valor float64) {
	b.valor = getValorBoleto(valor)
}

// SetNossoNumero - atribuir valor do nosso número (sem dv) boleto Bradesco
func (b *Bradesco) SetNossoNumero(nossoNumero string) {
	b.nossoNumero = getNossoNumero(nossoNumero, 11)
}

func (b *Bradesco) getDVCodigoBarras(codigoBarras CodigoBarras) string {

	codigo := b.retornarCodigoBarrasCompleto(codigoBarras)

	indice := len(codigo)
	multiplicador := 2
	var digito int
	var total int

	for indice >= 1 {
		digito, _ = strconv.Atoi(string(codigo[indice-1]))
		total += digito * multiplicador
		indice--
		multiplicador++
		if multiplicador > 9 {
			multiplicador = 2
		}
	}

	resultado := 11 - (total % 11)

	if resultado == 0 || resultado == 1 || resultado > 9 {
		return "1"
	}

	return strconv.Itoa(resultado)
}

func (b *Bradesco) getCampoLivre() string {
	return fmt.Sprintf("%s%s%s%s%d", b.agencia, b.carteira, b.nossoNumero, b.contaBeneficiario, 0)
}

func (b *Bradesco) retornarCodigoBarrasCompleto(cod CodigoBarras) string {

	return fmt.Sprintf("%s%s%s%s%s%s", cod.CodigoBanco, cod.CodigoMoeda, cod.DV, cod.FatorVencimento, cod.Valor, cod.CampoLivre)
}

// GetCodigoBarras - retorna objeto código de barras boleto Bradesco
func (b *Bradesco) GetCodigoBarras() CodigoBarras {

	cod := CodigoBarras{
		CodigoBanco:     b.codigoBradesco,
		CodigoMoeda:     "9",
		FatorVencimento: b.fatorVencimento,
		Valor:           b.valor,
		CampoLivre:      b.getCampoLivre(),
	}

	cod.DV = b.getDVCodigoBarras(cod)
	cod.CodigoBarrasCompleto = b.retornarCodigoBarrasCompleto(cod)

	return cod
}
