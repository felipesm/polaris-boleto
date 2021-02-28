package boletos

import (
	"fmt"
	"log"
	"strconv"

	errorutil "github.com/felipesm/polaris-boleto/erro"
)

// Bradesco - estrutura para representar o boleto Bradesco
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
func (b *Bradesco) SetCarteira(carteira int16) {
	b.carteira = fmt.Sprintf("%02d", carteira)
}

// SetCodigoBeneficiario - atribuir conta beneficiário boleto Bradesco
func (b *Bradesco) SetCodigoBeneficiario(contaBeneficiario int32) {
	b.contaBeneficiario = fmt.Sprintf("%07d", contaBeneficiario)
}

/*
SetFatorVencimento - atribuir fator de vencimento boleto Bradesco.
O "zerarVencimento" indica se o vencimento deve vir zerado no código de barras e linha digitável.
*/
func (b *Bradesco) SetFatorVencimento(dataVencimento string, zerarVencimento bool) {
	b.fatorVencimento = getFatorVencimento(dataVencimento, zerarVencimento)
}

/*
SetValorBoleto - atribuir valor boleto Bradesco.
O "zerarValor" indica se o valor do boleto deve vir zerado no código de barras e linha digitável.
*/
func (b *Bradesco) SetValorBoleto(valor float64, zerarValor bool) {
	b.valor = getValorBoleto(valor, zerarValor)
}

// SetNossoNumero - atribuir nosso número boleto Bradesco
func (b *Bradesco) SetNossoNumero(nossoNumero string) {
	b.nossoNumero = getNossoNumero(nossoNumero, 11)
}

func (b *Bradesco) getDVCodigoBarras(codigoBarras CodigoBarras) string {

	codigo := b.retornarCodigoBarrasCompleto(codigoBarras)
	total := getCalculoBaseCodigoBarras(codigo, 2, 9)
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

// GetCodigoBarras - retorna código de barras boleto Bradesco
func (b *Bradesco) GetCodigoBarras() (CodigoBarras, errorutil.Erro) {

	codigo := CodigoBarras{
		CodigoBanco:     b.codigoBradesco,
		CodigoMoeda:     "9",
		FatorVencimento: b.fatorVencimento,
		Valor:           b.valor,
		CampoLivre:      b.getCampoLivre(),
	}

	erro := b.validarDados()

	if erro.Status != 0 {
		codigo.DV = ""
		codigo.CodigoBarrasCompleto = ""
	} else {
		codigo.DV = b.getDVCodigoBarras(codigo)
		codigo.CodigoBarrasCompleto = b.retornarCodigoBarrasCompleto(codigo)
	}

	return codigo, erro
}

// GetLinhaDigitavel - retorna linha digitável boleto Bradesco
func (b *Bradesco) GetLinhaDigitavel(codigoBarras string) (LinhaDigitavel, errorutil.Erro) {

	var linha LinhaDigitavel
	var erro errorutil.Erro

	if len(codigoBarras) != 44 {
		erro = errorutil.CodigoBarrasInvalido(codigoBarras)
		log.Println(erro.Mensagem)
		return linha, erro
	}

	linha = getLinhaDigitavel(codigoBarras, 1, 2, true)
	return linha, erro
}

func (b *Bradesco) validarDados() errorutil.Erro {

	var erro errorutil.Erro

	if len(b.agencia) != 4 {
		erro = errorutil.AgenciaInvalida(b.agencia, 4)
		log.Println("Erro:", erro.Mensagem)
		return erro
	}

	if len(b.carteira) != 2 {
		erro = errorutil.CarteiraInvalida(b.carteira, 2)
		log.Println("Erro:", erro.Mensagem)
		return erro
	}

	if b.fatorVencimento == "0" {
		erro = errorutil.VencimentoInvalido()
		log.Println("Erro:", erro.Mensagem)
		return erro
	}

	if len(b.contaBeneficiario) != 7 {
		erro = errorutil.CodigoBeneficiarioInvalido(b.contaBeneficiario, 7)
		log.Println("Erro:", erro.Mensagem)
		return erro
	}

	if len(b.nossoNumero) != 11 {
		erro = errorutil.NossoNumeroInvalido(b.nossoNumero, 11)
		log.Println("Erro:", erro.Mensagem)
		return erro
	}

	return erro
}
