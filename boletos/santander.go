package boletos

import (
	"fmt"
	"log"
	"strconv"
)

// Santander - estrutura para representar o boleto Santander
type Santander struct {
	codigoSantander string // 01 a 03 - identificacao do banco
	codigoMoeda     string // 04 a 04 - codigo da moeda (real = 9 | outros = 0)
	dvCodigoBarras  string // 05 a 05 - digito verificador codigo de barras
	fatorVencimento string // 06 a 09 - fator vencimento
	valor           string // 10 a 19 - valor
	fixo1           string // 20 a 20 - fixo "9" (campo livre)
	codigoCedente   string // 21 a 27 - codigoCedente (campo livre)
	nossoNumero     string // 28 a 40 - nosso numero (campo livre)
	iosSeguradora   string // 41 a 41 - ios seguradora (preenchido com '0')
	carteira        string // 42 a 44 - modalidade carteira (101 = cobranca registro | 102 = cobranca s/ registro | 201 = penhor)
}

const codigoSantander string = "033"

// SetCodigo - atribuir código boleto Santander
func (s *Santander) SetCodigo() {
	s.codigoSantander = codigoSantander
}

// SetAgencia - não utilizado para boleto Santander
func (s *Santander) SetAgencia(agencia int32) {
}

// SetCarteira - atribuir código carteira boleto Santander
func (s *Santander) SetCarteira(carteira int16) {
	s.carteira = fmt.Sprintf("%03d", carteira)
}

// SetCodigoBeneficiario - atribuir código cedente boleto Santander
func (s *Santander) SetCodigoBeneficiario(codigoCedente int32) {
	s.codigoCedente = fmt.Sprintf("%07d", codigoCedente)
}

/*
SetFatorVencimento - atribuir fator de vencimento boleto Santander.
O "zerarVencimento" indica se o vencimento deve vir zerado no código de barras e linha digitável.
*/
func (s *Santander) SetFatorVencimento(dataVencimento string, zerarVencimento bool) {
	s.fatorVencimento = getFatorVencimento(dataVencimento, zerarVencimento)
}

/*
SetValorBoleto - atribuir valor boleto Santander.
O "zerarValor" indica se o valor do boleto deve vir zerado no código de barras e linha digitável.
*/
func (s *Santander) SetValorBoleto(valor float64, zerarValor bool) {
	s.valor = getValorBoleto(valor, zerarValor)
}

// SetNossoNumero - atribuir nosso número boleto Santander
func (s *Santander) SetNossoNumero(nossoNumero string) {
	s.nossoNumero = getNossoNumero(nossoNumero, 13)
}

func (s *Santander) getDVCodigoBarras(codigoBarras CodigoBarras) string {

	codigo := s.retornarCodigoBarrasCompleto(codigoBarras)

	total := getCalculoBaseCodigoBarras(codigo, 2, 9)

	resultado := ((total * 10) % 11)

	if resultado == 0 || resultado == 1 || resultado == 10 {
		return "1"
	}

	return strconv.Itoa(resultado)
}

func (s *Santander) getCampoLivre() string {
	return fmt.Sprintf("%d%s%s%d%s", 9, s.codigoCedente, s.nossoNumero, 0, s.carteira)
}

func (s *Santander) retornarCodigoBarrasCompleto(cod CodigoBarras) string {

	return fmt.Sprintf("%s%s%s%s%s%s", cod.CodigoBanco, cod.CodigoMoeda, cod.DV, cod.FatorVencimento, cod.Valor, cod.CampoLivre)
}

// GetCodigoBarras - retorna código de barras boleto Santander
func (s *Santander) GetCodigoBarras() (CodigoBarras, Erro) {

	codigo := CodigoBarras{
		CodigoBanco:     s.codigoSantander,
		CodigoMoeda:     "9",
		FatorVencimento: s.fatorVencimento,
		Valor:           s.valor,
		CampoLivre:      s.getCampoLivre(),
	}

	erro := s.validarDados()

	if erro.Status != 0 {
		codigo.DV = ""
		codigo.CodigoBarrasCompleto = ""
	} else {
		codigo.DV = s.getDVCodigoBarras(codigo)
		codigo.CodigoBarrasCompleto = s.retornarCodigoBarrasCompleto(codigo)
	}

	return codigo, erro
}

// GetLinhaDigitavel - retorna linha digitável boleto Santander
func (s *Santander) GetLinhaDigitavel(codigoBarras string) (LinhaDigitavel, Erro) {

	var linha LinhaDigitavel
	var erro Erro

	if len(codigoBarras) != 44 {
		erro = codigoBarrasInvalido(codigoBarras)
		log.Println(erro.Mensagem)
		return linha, erro
	}

	linha = getLinhaDigitavel(codigoBarras, 1, 2, true)
	return linha, erro
}

func (s *Santander) validarDados() Erro {

	var erro Erro

	if len(s.carteira) != 3 {
		erro = carteiraInvalida(s.carteira, 3)
		log.Println("Erro:", erro.Mensagem)
		return erro
	}

	if len(s.codigoCedente) != 7 {
		erro = codigoBeneficiarioInvalido(s.codigoCedente, 7)
		log.Println("Erro:", erro.Mensagem)
		return erro
	}

	if len(s.nossoNumero) != 13 {
		erro = nossoNumeroInvalido(s.nossoNumero, 13)
		log.Println("Erro:", erro.Mensagem)
		return erro
	}

	return erro
}
