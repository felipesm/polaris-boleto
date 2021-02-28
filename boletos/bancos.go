package boletos

import errorutil "github.com/felipesm/polaris-boleto/erro"

// Banco - interface que representa diferentes bancos
type Banco interface {
	SetCodigo()
	SetAgencia(agencia int32)
	SetCarteira(carteira int16)
	SetCodigoBeneficiario(contaBeneficiario int32)
	SetFatorVencimento(dataVencimento string, zerarVencimento bool)
	SetValorBoleto(valor float64, zerarValor bool)
	SetNossoNumero(nossoNumero string)
	getDVCodigoBarras(codigoBarras CodigoBarras) string
	getCampoLivre() string
	validarDados() errorutil.Erro
	GetCodigoBarras() (CodigoBarras, errorutil.Erro)
	GetLinhaDigitavel(codigoBarras string) (LinhaDigitavel, errorutil.Erro)
}
